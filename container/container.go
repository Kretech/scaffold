package container

import (
	"fmt"
)

type (
	Instance interface{}
)

type Container struct {
	//	记录别名
	aliases map[string]string

	//	标记是否已被解析过
	resolved map[string]bool

	//	记录所有绑定到容器的实体
	bindings map[string]entityAble

	//	这两个变量共同表示共享实体的信息
	//	前者表示是否是共享实体
	//	后者用来记录该实体解析之后的实例
	shares          map[string]bool
	sharedInstances map[string]Instance
}

func NewContainer() *Container {
	instance := &Container{}

	instance.init()

	return instance
}

func (c *Container) init() {
	c.aliases = make(map[string]string)
	c.bindings = make(map[string]entityAble)

	c.resolved = make(map[string]bool)

	c.shares = make(map[string]bool)
	c.sharedInstances = make(map[string]Instance)
}

func (c *Container) dropStaleInstance(abstract string) {
	c.bindings[abstract] = nil
	c.shares[abstract] = false
	c.aliases[abstract] = ""
}

func (c *Container) BindAny(abstract string, entity interface{}) {
	c.dropStaleInstance(abstract)

	concrete := c.toEntityAble(entity)

	c.bindings[abstract] = concrete
}

func (c *Container) Singleton(abstract string, entity interface{}) {
	c.BindAny(abstract, entity)
	c.share(abstract)
}

func (c *Container) share(abstract string) {
	c.shares[abstract] = true
}

//	Alias a type to a different name.
//	为 abstract 提供 alias 作为别名
func (c *Container) Alias(abstract string, alias string) {
	c.aliases[alias] = abstract
}

func (c *Container) Make(abstract string, params ...interface{}) interface{} {
	abstract = c.parseAlias(abstract)

	if c.sharedInstances[abstract] != nil {
		return c.sharedInstances[abstract]
	}

	concrete := c.getEntityAble(abstract)

	obj := c.buildEntity(concrete, params)

	if c.IsShared(abstract) {
		c.sharedInstances[abstract] = obj
	}

	return obj
}

func (c *Container) parseAlias(alias string) string {

	for c.aliases[alias] != `` {
		alias = c.aliases[alias]
	}

	return alias
}

func (c *Container) getEntityAble(abstract string) entityAble {

	binding := c.bindings[abstract]

	if binding == nil {
		panic(fmt.Sprintf(`unregisted binding [%s]`, abstract))
	}

	return binding
}

func (c *Container) buildEntity(binding entityAble, params []interface{}) interface{} {
	return binding.BuildEntity(c)
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

func (c *Container) toEntityAble(concrete interface{}) entityAble {
	switch concrete.(type) {

	case entityAble:
		return concrete.(entityAble)

	case funcEntity:
		return concrete.(funcEntity)

	case func() interface{}:
		return funcEntity(concrete.(func() interface{}))

	case funcWithSelfEntity:
		return concrete.(funcWithSelfEntity)

	case func(*Container) interface{}:
		return funcWithSelfEntity(concrete.(func(*Container) interface{}))

	default:
		return funcEntity(func() interface{} {
			return concrete
		})
	}
}

func (c *Container) BeforeResolving(string, interface{}) {
	panic("implement me")
}

func (c *Container) AfterResolved(string, interface{}) {
	panic("implement me")
}
