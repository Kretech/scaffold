package events

import (
	. "github.com/Kretech/scaffold/support"
)

type Dispatcher struct {
	events map[string][]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{make(map[string][]Listener)}
}

func (dispatcher *Dispatcher) ListenFunc(eventName string, fn interface{}) {

	dispatcher.initKey(eventName)

	listener := NewFuncListener(fn)

	dispatcher.events[eventName] = append(dispatcher.events[eventName], listener)
}

func (dispatcher *Dispatcher) Listen(event Event, listener Listener) {

	eventName := dispatcher.getEventKey(event)

	dispatcher.initKey(eventName)

	dispatcher.events[eventName] = append(dispatcher.events[eventName], listener)
}

func (dispatcher *Dispatcher) Dispatch(event Event) {

	var eventName string

	if name, ok := event.(string); ok {
		eventName = name
	} else {
		eventName = dispatcher.getEventKey(event)
	}

	dispatcher.initKey(eventName)

	if len(dispatcher.events[eventName]) < 1 {

	}

	for _, listener := range dispatcher.events[eventName] {
		listener.Handle(event)
	}
}

func (dispatcher *Dispatcher) getEventKey(event Event) string {
	return ClassName(event)
}

func (dispatcher *Dispatcher) initKey(eventName string) {
	if dispatcher.events[eventName] == nil {

	}
}
