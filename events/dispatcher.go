package events

type Dispatcher struct {
	events map[string][]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{make(map[string][]Listener)}
}

// ListenFunc 注册一个简单函数
//
// ListenFunc(e, func() {})
// ListenFunc(e, func() error {})
// ListenFunc(e, func(e))
// ListenFunc(e, func(e) error {})
// ListenFunc(e, foo.fn)
//
// @see scaffold/events/FuncListener.go#NewFuncListener
//
func (dispatcher *Dispatcher) ListenFunc(event Event, fn interface{}) {

	eventName := GetEventKey(event)

	listener := NewFuncListener(fn)

	dispatcher.events[eventName] = append(dispatcher.events[eventName], listener)
}

// Listen
func (dispatcher *Dispatcher) Listen(event Event, listener Listener) {

	eventName := GetEventKey(event)

	dispatcher.events[eventName] = append(dispatcher.events[eventName], listener)
}

// Dispatch 触发事件
func (dispatcher *Dispatcher) Dispatch(event Event) (err error) {

	eventName := GetEventKey(event)

	if len(dispatcher.events[eventName]) < 1 {
		return
	}

	for _, listener := range dispatcher.events[eventName] {
		err = listener.Handle(event)
		if err != nil {
			return
		}
	}

	return
}
