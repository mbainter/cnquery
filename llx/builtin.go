package llx

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.mondoo.io/mondoo/lumi"
	"go.mondoo.io/mondoo/types"
)

type chunkHandler struct {
	Compiler func(types.Type, types.Type) (string, error)
	f        func(*LeiseExecutor, *RawData, *Chunk, int32) (*RawData, int32, error)
	Label    string
}

// BuiltinFunctions for all builtin types
var BuiltinFunctions map[types.Type]map[string]chunkHandler

func init() {
	BuiltinFunctions = map[types.Type]map[string]chunkHandler{
		types.Bool: {
			string("==" + types.Bool):                {f: boolCmpBool, Label: "=="},
			string("!=" + types.Bool):                {f: boolNotBool, Label: "!="},
			string("==" + types.String):              {f: boolCmpString, Label: "=="},
			string("!=" + types.String):              {f: boolNotString, Label: "!="},
			string("==" + types.Regex):               {f: boolCmpRegex, Label: "=="},
			string("!=" + types.Regex):               {f: boolNotRegex, Label: "!="},
			string("==" + types.Array(types.Bool)):   {f: boolCmpBoolarray, Label: "=="},
			string("!=" + types.Array(types.Bool)):   {f: boolNotBoolarray, Label: "!="},
			string("==" + types.Array(types.String)): {f: boolCmpStringarray, Label: "=="},
			string("!=" + types.Array(types.String)): {f: boolNotStringarray, Label: "!="},
		},
		types.Int: {
			string("==" + types.Int):                 {f: intCmpInt, Label: "=="},
			string("!=" + types.Int):                 {f: intNotInt, Label: "!="},
			string("==" + types.String):              {f: intCmpString, Label: "=="},
			string("!=" + types.String):              {f: intNotString, Label: "!="},
			string("==" + types.Regex):               {f: intCmpRegex, Label: "=="},
			string("!=" + types.Regex):               {f: intNotRegex, Label: "!="},
			string("==" + types.Array(types.Int)):    {f: intCmpIntarray, Label: "=="},
			string("!=" + types.Array(types.Int)):    {f: intNotIntarray, Label: "!="},
			string("==" + types.Array(types.String)): {f: intCmpStringarray, Label: "=="},
			string("!=" + types.Array(types.String)): {f: intNotStringarray, Label: "!="},
			string("<" + types.Int):                  {f: intLTInt, Label: "<"},
			string("<=" + types.Int):                 {f: intLTEInt, Label: "<="},
			string(">" + types.Int):                  {f: intGTInt, Label: ">"},
			string(">=" + types.Int):                 {f: intGTEInt, Label: ">="},
			string("<" + types.Float):                {f: intLTFloat, Label: "<"},
			string("<=" + types.Float):               {f: intLTEFloat, Label: "<="},
			string(">" + types.Float):                {f: intGTFloat, Label: ">"},
			string(">=" + types.Float):               {f: intGTEFloat, Label: ">="},
			string("<" + types.String):               {f: intLTString, Label: "<"},
			string("<=" + types.String):              {f: intLTEString, Label: "<="},
			string(">" + types.String):               {f: intGTString, Label: ">"},
			string(">=" + types.String):              {f: intGTEString, Label: ">="},
		},
		types.Float: {
			string("==" + types.Float):               {f: floatCmpFloat, Label: "=="},
			string("!=" + types.Float):               {f: floatNotFloat, Label: "!="},
			string("==" + types.String):              {f: floatCmpString, Label: "=="},
			string("!=" + types.String):              {f: floatNotString, Label: "!="},
			string("==" + types.Regex):               {f: floatCmpRegex, Label: "=="},
			string("!=" + types.Regex):               {f: floatNotRegex, Label: "!="},
			string("==" + types.Array(types.Float)):  {f: floatCmpFloatarray, Label: "=="},
			string("!=" + types.Array(types.Float)):  {f: floatNotFloatarray, Label: "!="},
			string("==" + types.Array(types.String)): {f: floatCmpStringarray, Label: "=="},
			string("!=" + types.Array(types.String)): {f: floatNotStringarray, Label: "!="},
			string("<" + types.Float):                {f: floatLTFloat, Label: "<"},
			string("<=" + types.Float):               {f: floatLTEFloat, Label: "<="},
			string(">" + types.Float):                {f: floatGTFloat, Label: ">"},
			string(">=" + types.Float):               {f: floatGTEFloat, Label: ">="},
			string("<" + types.Int):                  {f: floatLTInt, Label: "<"},
			string("<=" + types.Int):                 {f: floatLTEInt, Label: "<="},
			string(">" + types.Int):                  {f: floatGTInt, Label: ">"},
			string(">=" + types.Int):                 {f: floatGTEInt, Label: ">="},
			string("<" + types.String):               {f: floatLTString, Label: "<"},
			string("<=" + types.String):              {f: floatLTEString, Label: "<="},
			string(">" + types.String):               {f: floatGTString, Label: ">"},
			string(">=" + types.String):              {f: floatGTEString, Label: ">="},
		},
		types.String: {
			string("==" + types.String):              {f: stringCmpString, Label: "=="},
			string("!=" + types.String):              {f: stringNotString, Label: "!="},
			string("==" + types.Regex):               {f: stringCmpRegex, Label: "=="},
			string("!=" + types.Regex):               {f: stringNotRegex, Label: "!="},
			string("==" + types.Bool):                {f: stringCmpBool, Label: "=="},
			string("!=" + types.Bool):                {f: stringNotBool, Label: "!="},
			string("==" + types.Int):                 {f: stringCmpInt, Label: "=="},
			string("!=" + types.Int):                 {f: stringNotInt, Label: "!="},
			string("==" + types.Float):               {f: stringCmpFloat, Label: "=="},
			string("!=" + types.Float):               {f: stringNotFloat, Label: "!="},
			string("==" + types.Array(types.String)): {f: stringCmpStringarray, Label: "=="},
			string("!=" + types.Array(types.String)): {f: stringNotStringarray, Label: "!="},
			string("==" + types.Array(types.Bool)):   {f: stringCmpBoolarray, Label: "=="},
			string("!=" + types.Array(types.Bool)):   {f: stringNotBoolarray, Label: "!="},
			string("==" + types.Array(types.Int)):    {f: stringCmpIntarray, Label: "=="},
			string("!=" + types.Array(types.Int)):    {f: stringNotIntarray, Label: "!="},
			string("==" + types.Array(types.Float)):  {f: stringCmpFloatarray, Label: "=="},
			string("!=" + types.Array(types.Float)):  {f: stringNotFloatarray, Label: "!="},
			string("<" + types.String):               {f: stringLTString, Label: "<"},
			string("<=" + types.String):              {f: stringLTEString, Label: "<="},
			string(">" + types.String):               {f: stringGTString, Label: ">"},
			string(">=" + types.String):              {f: stringGTEString, Label: ">="},
			string("<" + types.Int):                  {f: stringLTInt, Label: "<"},
			string("<=" + types.Int):                 {f: stringLTEInt, Label: "<="},
			string(">" + types.Int):                  {f: stringGTInt, Label: ">"},
			string(">=" + types.Int):                 {f: stringGTEInt, Label: ">="},
			string("<" + types.Float):                {f: stringLTFloat, Label: "<"},
			string("<=" + types.Float):               {f: stringLTEFloat, Label: "<="},
			string(">" + types.Float):                {f: stringGTFloat, Label: ">"},
			string(">=" + types.Float):               {f: stringGTEFloat, Label: ">="},
		},
		types.Regex: {
			string("==" + types.Regex):               {f: stringCmpString, Label: "=="},
			string("!=" + types.Regex):               {f: stringNotString, Label: "!="},
			string("==" + types.Bool):                {f: regexCmpBool, Label: "=="},
			string("!=" + types.Bool):                {f: regexNotBool, Label: "!="},
			string("==" + types.Int):                 {f: regexCmpInt, Label: "=="},
			string("!=" + types.Int):                 {f: regexNotInt, Label: "!="},
			string("==" + types.Float):               {f: regexCmpFloat, Label: "=="},
			string("!=" + types.Float):               {f: regexNotFloat, Label: "!="},
			string("==" + types.String):              {f: regexCmpString, Label: "=="},
			string("!=" + types.String):              {f: regexNotString, Label: "!="},
			string("==" + types.Array(types.Regex)):  {f: stringCmpStringarray, Label: "=="},
			string("!=" + types.Array(types.Regex)):  {f: stringNotStringarray, Label: "!="},
			string("==" + types.Array(types.Bool)):   {f: regexCmpBoolarray, Label: "=="},
			string("!=" + types.Array(types.Bool)):   {f: regexNotBoolarray, Label: "!="},
			string("==" + types.Array(types.Int)):    {f: regexCmpIntarray, Label: "=="},
			string("!=" + types.Array(types.Int)):    {f: regexNotIntarray, Label: "!="},
			string("==" + types.Array(types.Float)):  {f: regexCmpFloatarray, Label: "=="},
			string("!=" + types.Array(types.Float)):  {f: regexNotFloatarray, Label: "!="},
			string("==" + types.Array(types.String)): {f: regexCmpStringarray, Label: "=="},
			string("!=" + types.Array(types.String)): {f: regexNotStringarray, Label: "!="},
		},
		types.ArrayLike: {
			"[]":     {f: arrayGetIndex},
			"{}":     {f: arrayBlockList},
			"length": {f: arrayLength},
			"==":     {Compiler: compileArrayOpArray("==")},
			"!=":     {Compiler: compileArrayOpArray("!=")},
			// []T -- []T
			string(types.Bool + "==" + types.Array(types.Bool)):     {f: boolarrayCmpBoolarray, Label: "=="},
			string(types.Bool + "!=" + types.Array(types.Bool)):     {f: boolarrayNotBoolarray, Label: "!="},
			string(types.Int + "==" + types.Array(types.Int)):       {f: intarrayCmpIntarray, Label: "=="},
			string(types.Int + "!=" + types.Array(types.Int)):       {f: intarrayNotIntarray, Label: "!="},
			string(types.Float + "==" + types.Array(types.Float)):   {f: floatarrayCmpFloatarray, Label: "=="},
			string(types.Float + "!=" + types.Array(types.Float)):   {f: floatarrayNotFloatarray, Label: "!="},
			string(types.String + "==" + types.Array(types.String)): {f: stringarrayCmpStringarray, Label: "=="},
			string(types.String + "!=" + types.Array(types.String)): {f: stringarrayNotStringarray, Label: "!="},
			string(types.Regex + "==" + types.Array(types.Regex)):   {f: stringarrayCmpStringarray, Label: "=="},
			string(types.Regex + "!=" + types.Array(types.Regex)):   {f: stringarrayNotStringarray, Label: "!="},
			// []T -- T
			string(types.Bool + "==" + types.Bool):     {f: boolarrayCmpBool, Label: "=="},
			string(types.Bool + "!=" + types.Bool):     {f: boolarrayNotBool, Label: "!="},
			string(types.Int + "==" + types.Int):       {f: intarrayCmpInt, Label: "=="},
			string(types.Int + "!=" + types.Int):       {f: intarrayNotInt, Label: "!="},
			string(types.Float + "==" + types.Float):   {f: floatarrayCmpFloat, Label: "=="},
			string(types.Float + "!=" + types.Float):   {f: floatarrayNotFloat, Label: "!="},
			string(types.String + "==" + types.String): {f: stringarrayCmpString, Label: "=="},
			string(types.String + "!=" + types.String): {f: stringarrayNotString, Label: "!="},
			string(types.Regex + "==" + types.Regex):   {f: stringarrayCmpString, Label: "=="},
			string(types.Regex + "!=" + types.Regex):   {f: stringarrayNotString, Label: "!="},
			// []string -- T
			string(types.String + "==" + types.Bool):  {f: stringarrayCmpBool, Label: "=="},
			string(types.String + "!=" + types.Bool):  {f: stringarrayNotBool, Label: "!="},
			string(types.String + "==" + types.Int):   {f: stringarrayCmpInt, Label: "=="},
			string(types.String + "!=" + types.Int):   {f: stringarrayNotInt, Label: "!="},
			string(types.String + "==" + types.Float): {f: stringarrayCmpFloat, Label: "=="},
			string(types.String + "!=" + types.Float): {f: stringarrayNotFloat, Label: "!="},
			// []T -- string
			string(types.Bool + "==" + types.String):  {f: boolarrayCmpString, Label: "=="},
			string(types.Bool + "!=" + types.String):  {f: boolarrayNotString, Label: "!="},
			string(types.Int + "==" + types.String):   {f: intarrayCmpString, Label: "=="},
			string(types.Int + "!=" + types.String):   {f: intarrayNotString, Label: "!="},
			string(types.Float + "==" + types.String): {f: floatarrayCmpString, Label: "=="},
			string(types.Float + "!=" + types.String): {f: floatarrayNotString, Label: "!="},
			// []T -- regex
			string(types.Bool + "==" + types.Regex):   {f: boolarrayCmpRegex, Label: "=="},
			string(types.Bool + "!=" + types.Regex):   {f: boolarrayNotRegex, Label: "!="},
			string(types.Int + "==" + types.Regex):    {f: intarrayCmpRegex, Label: "=="},
			string(types.Int + "!=" + types.Regex):    {f: intarrayNotRegex, Label: "!="},
			string(types.Float + "==" + types.Regex):  {f: floatarrayCmpRegex, Label: "=="},
			string(types.Float + "!=" + types.Regex):  {f: floatarrayNotRegex, Label: "!="},
			string(types.String + "==" + types.Regex): {f: stringarrayCmpRegex, Label: "=="},
			string(types.String + "!=" + types.Regex): {f: stringarrayNotRegex, Label: "!="},
		},
		types.MapLike: {
			"[]":     {f: mapGetIndex},
			"length": {f: mapLength},
		},
		types.ResourceLike: {
			"where":  {f: resourceWhere},
			"length": {f: resourceLength},
			"{}": {f: func(c *LeiseExecutor, bind *RawData, chunk *Chunk, ref int32) (*RawData, int32, error) {
				return c.runBlock(bind, chunk.Function.Args[0], ref)
			}},
		},
	}
}

func runResourceFunction(c *LeiseExecutor, bind *RawData, chunk *Chunk, ref int32) (*RawData, int32, error) {
	// ugh something is wrong here.... fix it later
	rr, ok := bind.Value.(lumi.ResourceType)
	if !ok {
		// TODO: can we get rid of this fmt call
		return nil, 0, fmt.Errorf("cannot cast resource to resource type: %+v", bind.Value)
	}

	info := rr.LumiResource()
	// resource := c.runtime.Registry.Resources[bind.Type]
	if info == nil {
		return nil, 0, errors.New("Cannot retrieve resource from the binding to run the raw function")
	}

	resource, ok := c.runtime.Registry.Resources[info.Name]
	if !ok || resource == nil {
		return nil, 0, errors.New("Cannot retrieve resource definition for resource '" + info.Name + "'")
	}

	// record this watcher on the executors watcher IDs
	wid := c.watcherUID(ref)
	log.Debug().Str("wid", wid).Msg("exec> add watcher id ")
	c.watcherIds.Store(wid)

	// watch this field in the resource
	err := c.runtime.WatchAndUpdate(rr, chunk.Id, wid, func(fieldData interface{}, fieldError error) {
		if fieldError != nil {
			c.callback(errorResult(fieldError, c.entrypoints[ref]))
			return
		}

		c.cache.Store(ref, &stepCache{Result: &RawData{
			Type:  types.Type(resource.Fields[chunk.Id].Type),
			Value: fieldData,
			Error: fieldError,
		}})
		c.triggerChain(ref)
	})

	// we are done executing this chain
	return nil, 0, err
}

// BuiltinFunction provides the handler for this type's function
func BuiltinFunction(typ types.Type, name string) (*chunkHandler, error) {
	h, ok := BuiltinFunctions[typ.Underlying()]
	if !ok {
		return nil, errors.New("cannot find functions for type '" + typ.Label() + "' (called '" + name + "')")
	}
	fh, ok := h[name]
	if !ok {
		return nil, errors.New("cannot find function '" + name + "' for type '" + typ.Label() + "'")
	}
	return &fh, nil
}

// this is called for objects that call a function
func (c *LeiseExecutor) runBoundFunction(bind *RawData, chunk *Chunk, ref int32) (*RawData, int32, error) {
	log.Debug().Int32("ref", ref).Str("id", chunk.Id).Msg("exec> run bound function")

	fh, err := BuiltinFunction(bind.Type, chunk.Id)
	if err == nil {
		res, dref, err := fh.f(c, bind, chunk, ref)
		if res != nil {
			c.cache.Store(ref, &stepCache{Result: res})
		}
		return res, dref, err
	}

	if bind.Type.IsResource() {
		return runResourceFunction(c, bind, chunk, ref)
	}
	return nil, 0, err
}
