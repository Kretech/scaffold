package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIniRepository(t *testing.T) {

	Convey(
		"TestRepository",
		t,
		func() {

			c := NewRepository()
			c.Set(`app.mode`, `dev`)

			So(c.Get(`app.mode`) == `dev`, ShouldBeTrue)
			So(c.Get(`app.mode`), ShouldEqual, `dev`)

		},
	)

}
