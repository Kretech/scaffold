package events

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDispatcher_ListenFunc(t *testing.T) {

	Convey("TestDispatcherListenFunc", t, func() {
		dispatcher := NewDispatcher()

		start := false

		dispatcher.ListenFunc(`start`, func() {
			start = true
		})

		dispatcher.Dispatch(`start`)

		So(start, ShouldBeTrue)
	})
}

type FakeProcess struct{ State int }
type FakeListener1 struct{}

func (Listener *FakeListener1) Handle(event Event) (err error) {
	event.(*FakeProcess).State += 1
	return
}

type FakeListener2 struct{}

func (Listener *FakeListener2) Handle(event Event) (err error) {
	event.(*FakeProcess).State += 2
	return
}

func TestDispatcher_ListenEvent(t *testing.T) {

	Convey("Listen", t, func() {

		p := &FakeProcess{}

		dispatcher := NewDispatcher()

		dispatcher.Listen(p, &FakeListener1{})
		dispatcher.Listen(p, &FakeListener2{})

		So(p.State, ShouldEqual, 0)

		dispatcher.Dispatch(p)
		So(p.State, ShouldEqual, 3)
	})
}
