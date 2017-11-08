package container

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContainer_BindFunc(t *testing.T) {

	data := [][]interface{}{{
		"name", func() interface{} {
			return `huiren`
		},
		`huiren`,
	}}

	Convey(
		"TestContainerBindFunc",
		t,
		func() {
			c := NewContainer()

			for i := range data {

				abstract := data[i][0].(string)
				binding := data[i][1].(func() interface{})
				expect := data[i][2]

				c.BindFunc(abstract, binding)

				result := c.Make(abstract)

				So(result, ShouldEqual, expect)
			}
		},
	)
}

func TestContainer_Singleton(t *testing.T) {

	Convey(
		"TestSingleton",
		t,
		func() {
			c := NewContainer()

			c.Singleton(`notRandom`, func(c *Container) interface{} {
				return rand.Int63()
			})

			So(c.Make(`notRandom`) == c.Make(`notRandom`), ShouldBeTrue)
		},
	)
}

func TestContainer_Alias(t *testing.T) {

	Convey(
		"TestAlias",
		t,
		func() {
			c := NewContainer()

			c.BindAny(`name`, func(c *Container) interface{} {
				return `huiren`
			})

			c.Alias(`name`, `alias_name`)

			So(c.Make(`name`) == c.Make(`alias_name`), ShouldBeTrue)
		},
	)
}

func TestContainer_BindInstance(t *testing.T) {
	Convey(
		"TestContainerBindInstance",
		t,
		func() {
			type Person struct {
				Name string
				Age  int
			}

			p := new(Person)
			p.Name = `test`
			p.Age = 12

			c := NewContainer()

			c.BindInstance(`person`, p)
			So(c.Make(`person`) == p, ShouldBeTrue)

			c.BindInstance(`person`, *p)
			So(c.Make(`person`) == *p, ShouldBeTrue)
		},
	)
}
