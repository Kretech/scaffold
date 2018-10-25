package events

type FuncListener struct {
	fn func(Event) error
}

//
// func() {}
// func() error {}
// func(e)
// func(e) error {}
// foo.fn
//
func NewFuncListener(fn interface{}) *FuncListener {

	var fe func(Event) error

	switch f := fn.(type) {
	case func():
		fe = func(e Event) (err error) {
			f()
			return
		}

	case func() error:
		fe = func(e Event) error {
			return f()
		}

	case func(e Event):
		fe = func(e Event) (err error) {
			f(e)
			return
		}

	case func(e Event) error:
		fe = f

	default:
		panic("nonsupport func listener")

	}

	return &FuncListener{fe}
}

func (fl *FuncListener) Handle(event Event) (err error) {
	return fl.fn(event)
}
