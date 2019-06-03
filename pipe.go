package fnpipe

import (
	"fmt"
	"reflect"
)

type Pipe struct {
	ls []interface{}
}

func NewPipeline(ls ...interface{}) (*Pipe, error) {
	var pipe []interface{}

	p := &Pipe{
		ls: pipe,
	}

	for _, f := range ls {
		if err := p.Add(f); err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *Pipe) Add(pf interface{}) error {
	if reflect.ValueOf(pf).Type().Kind() != reflect.Func {
		return fmt.Errorf("pipe func should be type Func")
	}

	p.ls = append(p.ls, pf)

	return nil
}

func (p *Pipe) ExecWith(input ...interface{}) (error, interface{}) {
	var output []interface{}

	for i, e := range p.ls {
		output = make([]interface{}, 0)
		fn := reflect.ValueOf(e)
		if fn.Type().NumIn() != len(input) {
			// TODO: check kind of each input to match fn definition
			return fmt.Errorf("pipe: argument mismatch in pipeline func #%d", i), nil
		}

		// build []reflect.Value for fn input
		val := make([]reflect.Value, 0)
		for _, in := range input {
			val = append(val, reflect.ValueOf(in))
		}

		// call the func
		o := fn.Call(val)

		for _, u := range o {
			// TODO: check kind of each output to match fn definition
			output = append(output, reflect.Indirect(u).Interface())
		}

		// pass this output to the coming next function
		// if it was the last fn, we will just extract data from []reflect.Value
		// in order to have concrete return value
		input = output
	}

	values := make([]interface{}, 0)
	for _, o := range output {
		values = append(values, reflect.ValueOf(o).Interface())
	}

	// we ensure the 1st piece of the result is always the error
	return nil, values
}