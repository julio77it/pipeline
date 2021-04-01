package pipeline

import (
	"reflect"
)

// PipeLine : the event pipeline, one goroutine for work
// Arguments of the next work function have to be compatible
// with the result set of the previous one
type PipeLine struct {
	works []interface{}
	in    chan<- []interface{}
	out   <-chan []interface{}
}

// Process let the pipeline does the work
func (pl *PipeLine) Process(args ...interface{}) {
	pl.in <- args
}

// Chain : create a new PipeLine, sequence dependent works
func Chain(wks ...interface{}) (*PipeLine, error) {
	p := PipeLine{
		works: make([]interface{}, len(wks), len(wks)),
	}

	var in chan []interface{} = make(chan []interface{}, 0)
	var out chan []interface{} = make(chan []interface{}, 0)
	p.in = in

	for i, w := range wks {
		p.works[i] = w

		if i < len(wks) {
			out = make(chan []interface{}, 0)
		} else {
			out = nil
		}
		// run adapter
		go adapter(in, out, w)

		if i < len(wks) {
			in = out
		}
	}
	return &p, nil
}

func adapter(in <-chan []interface{}, out chan<- []interface{}, f interface{}) {
	val := reflect.ValueOf(f)
	typ := val.Type()

	valueArgs := make([]reflect.Value, typ.NumIn(), typ.NumIn())

	// get args from previous chained task
	args := <-in

	// convert []args in []reflect.Value for reflect.Call
	for i := 0; i < len(valueArgs); i++ {
		if i == len(args) { // not enough args, leave other to zero values
			break
		}
		valueArgs[i] = reflect.ValueOf(args[i])
	}
	// invoke function call
	valueResult := val.Call(valueArgs)

	// if not last chained task
	if out != nil {
		result := make([]interface{}, 0, len(valueResult))
		// convert []results to []reflect.Value for forward as arguments
		// to next chained task
		for _, value := range valueResult {
			result = append(result, value.Interface())
		}
		// forward arguments to next task
		out <- result
	}
}
