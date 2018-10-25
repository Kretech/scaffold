package container

type hasAlias struct {
	//	记录别名
	aliases map[string]string
}

//	Alias a type to a different name.
//	为 abstract 提供 alias 作为别名
func (c *hasAlias) Alias(abstract string, alias string) {
	c.aliases[alias] = abstract
}

func (c *hasAlias) parseAlias(alias string) string {

	for c.aliases[alias] != `` {
		alias = c.aliases[alias]
	}

	return alias
}
