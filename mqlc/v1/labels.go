package v1

import (
	"errors"
	"strconv"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"go.mondoo.io/mondoo/llx"
	"go.mondoo.io/mondoo/resources"
	"go.mondoo.io/mondoo/types"
)

func createArgLabel(arg *llx.Primitive, code *llx.CodeV1, labels *llx.Labels, schema *resources.Schema) error {
	if !types.Type(arg.Type).IsFunction() {
		return nil
	}

	ref, ok := arg.RefV1()
	if !ok {
		return errors.New("cannot get function reference")
	}

	function := code.Functions[ref-1]
	return UpdateLabels(function, labels, schema)
}

func createLabel(code *llx.CodeV1, ref int32, labels *llx.Labels, schema *resources.Schema) (string, error) {
	chunk := code.Code[ref-1]

	if chunk.Call == llx.Chunk_PRIMITIVE {
		return "", nil
	}

	id := chunk.Id
	if chunk.Function == nil {
		return id, nil
	}

	// TODO: workaround to get past the builtin global call
	// this needs proper handling for global calls
	if chunk.Function.DeprecatedV5Binding == 0 && id != "if" {
		return id, nil
	}

	var parentLabel string
	if chunk.Function.DeprecatedV5Binding != 0 {
		var err error
		parentLabel, err = createLabel(code, chunk.Function.DeprecatedV5Binding, labels, schema)
		if err != nil {
			return "", err
		}
	}

	var res string
	switch id {
	case "[]":
		if len(chunk.Function.Args) != 1 {
			panic("don't know how to extract label data from array access without args")
		}

		arg := chunk.Function.Args[0].RawData()
		idx := arg.Value

		switch arg.Type {
		case types.Int:
			res = "[" + strconv.FormatInt(idx.(int64), 10) + "]"
		case types.String:
			res = "[" + idx.(string) + "]"
		default:
			panic("cannot label array index of type " + arg.Type.Label())
		}
		if parentLabel != "" {
			res = parentLabel + res
		}
	case "{}", "${}":
		res = parentLabel

		fref := chunk.Function.Args[0]
		if !types.Type(fref.Type).IsFunction() {
			panic("don't know how to extract label data when argument is not a function: " + types.Type(fref.Type).Label())
		}

		ref, ok := fref.RefV1()
		if !ok {
			panic("cannot find function reference for data extraction")
		}

		function := code.Functions[ref-1]
		err := UpdateLabels(function, labels, schema)
		if err != nil {
			return "", err
		}

	case "if":
		res = "if"

		var i int
		max := len(chunk.Function.Args)
		for i+1 < max {
			arg := chunk.Function.Args[i+1]
			if err := createArgLabel(arg, code, labels, schema); err != nil {
				return "", err
			}
			i += 2
		}

		if i < max {
			arg := chunk.Function.Args[i]
			if err := createArgLabel(arg, code, labels, schema); err != nil {
				return "", err
			}
		}

	default:
		if label, ok := llx.ComparableLabel(id); ok {
			arg := chunk.Function.Args[0].LabelV1(code)
			res = parentLabel + " " + label + " " + arg
		} else if parentLabel == "" {
			res = id
		} else {
			res = parentLabel + "." + id
		}
	}

	// TODO: figure out why this string includes control characters in the first place
	return stripCtlAndExtFromUnicode(res), nil
}

// Unicode normalization and filtering, see http://blog.golang.org/normalization and
// http://godoc.org/golang.org/x/text/unicode/norm for more details.
func stripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	str, _, _ = transform.String(t, str)
	return str
}

// UpdateLabels for the given code under the schema
func UpdateLabels(code *llx.CodeV1, labels *llx.Labels, schema *resources.Schema) error {
	if code == nil {
		return errors.New("cannot create labels without code")
	}

	datapoints := code.Datapoints

	// We don't want assertions to become labels. Their data should not be printed
	// regularly but instead be processed through the assertion itself
	if code.Assertions != nil {
		assertionPoints := map[int32]struct{}{}
		for _, assertion := range code.Assertions {
			for j := range assertion.DeprecatedV5Datapoint {
				assertionPoints[assertion.DeprecatedV5Datapoint[j]] = struct{}{}
			}
		}

		filtered := []int32{}
		for i := range datapoints {
			ref := datapoints[i]
			if _, ok := assertionPoints[ref]; ok {
				continue
			}
			filtered = append(filtered, ref)
		}
		datapoints = filtered
	}

	labelrefs := append(code.Entrypoints, datapoints...)

	var err error
	for _, entrypoint := range labelrefs {
		checksum, ok := code.Checksums[entrypoint]
		if !ok {
			return errors.New("failed to create labels, cannot find checksum for this entrypoint " + strconv.FormatUint(uint64(entrypoint), 10))
		}

		if _, ok := labels.Labels[checksum]; ok {
			continue
		}

		labels.Labels[checksum], err = createLabel(code, entrypoint, labels, schema)
		if err != nil {
			return err
		}
	}

	// any more checksums that might have been set need to be removed, since we don't need them
	// TODO: there must be a way to do this without having to create the label first
	if code.Assertions != nil {
		for _, assertion := range code.Assertions {
			if !assertion.DecodeBlock {
				continue
			}

			for i := 0; i < len(assertion.Checksums); i++ {
				delete(labels.Labels, assertion.Checksums[i])
			}
		}
	}

	return nil
}