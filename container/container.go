package container

import (
	"fmt"
	"log"

	contractContainer "github.com/Kretech/contracts/container"
)

type (
	Instance interface{}
)

type Options struct {
	EnableLog bool
}

type Container struct {
	//	标记是否已被解析过
	resolved map[string]bool

	hasBinding
	hasAlias
	hasShared

	Options
}

func NewContainer() *Container {
	instance := &Container{}

	instance.init()

	instance.Options = Options{
		EnableLog: true,
	}

	return instance
}

func (c *Container) init() {
	c.aliases = make(map[string]string)
	c.bindings = make(map[string]entityAble)

	c.resolved = make(map[string]bool)

	c.hasShared.init()
}

func (c *Container) dropExisted(abstract string) {
	c.bindings[abstract] = nil
	c.shares[abstract] = false
	c.aliases[abstract] = ""
}

func (c *Container) BindAny(abstract string, entity interface{}) {
	c.dropExisted(abstract)

	concrete := newEntityAble(entity)

	c.bindings[abstract] = concrete
}

func (c *Container) Singleton(abstract string, entity interface{}) {
	c.BindAny(abstract, entity)
	c.markAsShared(abstract)
}

func (c *Container) Make(abstract string, params ...interface{}) interface{} {
	abstract = c.parseAlias(abstract)

	if c.Options.EnableLog {
		log.Println(`making:`, abstract)
		defer log.Println(`maked:`, abstract)
	}

	if c.sharedInstances[abstract] != nil {
		return c.sharedInstances[abstract]
	}

	concrete := c.getEntityAble(abstract)

	obj := buildEntity(concrete, params)
	if initer, ok := obj.(contractContainer.Initer); ok {
		initer.Init()
	}

	if c.IsShared(abstract) {
		c.sharedInstances[abstract] = obj
	}

	return obj
}

func (c *Container) getEntityAble(abstract string) entityAble {

	binding := c.bindings[abstract]

	if binding == nil {
		panic(fmt.Sprintf(`unregisted binding [%s]`, abstract))
	}

	return binding
}

func (c *Container) BindFunc(abstract string, fn func() interface{}) {
	c.BindAny(abstract, funcEntity(fn))
}

func (c *Container) BindInstance(abstract string, instance interface{}) {
	c.BindAny(abstract, instance)
}

func (c *Container) IsBound(abstract string) bool {
	return c.bindings[abstract] != nil
}

func (c *Container) IsResolved(abstract string) bool {
	return c.resolved[abstract]
}

func (c *Container) IsShared(abstract string) bool {
	return c.shares[abstract]
}

func (c *Container) IsAlias(alias string) bool {
	return len(c.aliases[alias]) > 0
}

func (c *Container) BeforeResolving(string, interface{}) {
	panic("implement me")
}

func (c *Container) AfterResolved(string, interface{}) {
	panic("implement me")
}
