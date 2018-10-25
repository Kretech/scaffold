package events

import (
	"github.com/Kretech/contracts/event"
	"github.com/Kretech/scaffold/support"
)

type Event = event.Event

func GetEventKey(evt Event) event.EventKey {

	switch t := evt.(type) {
	case string:
		return t

	case event.EventKeyHolder:
		return t.EventKey()

	}

	return support.ClassName(evt)
}
