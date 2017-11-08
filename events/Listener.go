package events

type Listener interface {
	Handle(Event)
}
