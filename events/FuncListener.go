package events

import "reflect"

type FuncListener struct {
	fn interface{}
}

func NewFuncListener(fn interface{}) *FuncListener {
	return &FuncListener{fn}
}

func (fl *FuncListener) Handle(event Event) {
	e := reflect.ValueOf(event)
	v := reflect.ValueOf(fl.fn)

	in := []reflect.Value{}
	if v.Type().NumIn() > 0 {
		in = append(in, e)
	}

	v.Call(in)
}
