package config

import (
	"github.com/Kretech/contracts/config"
	gotoml "github.com/pelletier/go-toml"
)

type Toml struct {
	*gotoml.Tree
}

func (t *Toml) GetSub(key string) config.Repository {
	return NewToml(t.Tree.Get(key).(*gotoml.Tree))
}

func NewToml(t *gotoml.Tree) *Toml {
	return &Toml{t}
}

func (t *Toml) Load(content []byte) error {
	tree, err := gotoml.LoadBytes(content)
	t.Tree = tree
	return err
}
