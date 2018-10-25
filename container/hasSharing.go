package container

type hasShared struct {
	//	这两个变量共同表示共享实体的信息
	//	前者表示是否是共享实体
	//	后者用来记录该实体解析之后的实例
	shares          map[string]bool
	sharedInstances map[string]Instance
}

func (c *hasShared) init() {
	c.shares = make(map[string]bool)
	c.sharedInstances = make(map[string]Instance)
}

func (c *hasShared) markAsShared(abstract string) {
	c.shares[abstract] = true
}
