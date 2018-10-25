package config

import gotoml "github.com/pelletier/go-toml"

type Toml struct {
	*gotoml.Tree
}

func (t *Toml) Load(content []byte) error {
	tree, err := gotoml.LoadBytes(content)
	t.Tree = tree
	return err
}
