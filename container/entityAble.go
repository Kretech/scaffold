package container

type entityAble interface {
	//	实例化实现
	BuildEntity(*Container, ...interface{}) interface{}
}

type funcEntity func() interface{}

func (e funcEntity) BuildEntity(c *Container, params ...interface{}) interface{} {
	return e()
}

type funcWithSelfEntity func(*Container) interface{}

func (e funcWithSelfEntity) BuildEntity(c *Container, params ...interface{}) interface{} {
	return e(c)
}
