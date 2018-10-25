package container

type funcEntity func() interface{}

func (e funcEntity) buildEmptyEntity(c *Container, params ...interface{}) interface{} {
	return e()
}

type funcWithSelfEntity func(*Container) interface{}

func (e funcWithSelfEntity) buildEmptyEntity(c *Container, params ...interface{}) interface{} {
	return e(c)
}
