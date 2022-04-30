package llx

import (
	"sync"

	"go.mondoo.io/mondoo/lumi"
)

// Run a piece of compiled code against a runtime. Just a friendly helper method
func RunV1(code *CodeV1, runtime *lumi.Runtime, props map[string]*Primitive, callback ResultCallback) (*LeiseExecutorV1, error) {
	x, err := NewExecutorV1(code, runtime, props, callback)
	if err != nil {
		return nil, err
	}
	x.Run()
	return x, nil
}

func NoRunV1(code *CodeV1, runerr error, runtime *lumi.Runtime, props map[string]*Primitive, callback ResultCallback) (*LeiseExecutorV1, error) {
	x, err := NewExecutorV1(code, runtime, props, callback)
	if err != nil {
		return nil, err
	}
	x.NoRun(runerr)
	return x, nil
}

// RunOnce the code that was provided and call the callback
func RunOnceV1(code *CodeV1, runtime *lumi.Runtime, props map[string]*Primitive, callback func(one *RawResult, isDone bool)) error {
	cnt := 0
	var executor *LeiseExecutorV1
	var err error

	maxCnt := len(code.Entrypoints) + len(code.Datapoints)

	// Note: We cannot copy the code from the Run method above as it may
	// lead to a race condition where the callback is run BEFORE the
	// executor is created. The way we do it here guarantees everything
	// including the closure-based executor is in place before the callback
	// runs.
	executor, err = NewExecutorV1(code, runtime, props, func(one *RawResult) {
		var isDone = false
		cnt++

		if cnt >= maxCnt {
			isDone = true
			executor.Unregister()
		}

		callback(one, isDone)
	})
	if err != nil {
		return err
	}

	executor.Run()
	return nil
}

// RunOnceSync will run the code only once and report on the results it gets
func RunOnceSyncV1(code *CodeV1, runtime *lumi.Runtime, props map[string]*Primitive) ([]*RawResult, error) {
	res := []*RawResult{}

	var done sync.WaitGroup
	done.Add(1)

	err := RunOnceV1(code, runtime, props, func(one *RawResult, isDone bool) {
		res = append(res, one)
		if isDone {
			done.Done()
		}
	})
	if err != nil {
		return nil, err
	}

	done.Wait()

	return res, nil
}