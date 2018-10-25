package support

import (
	"reflect"
)

type classLoader struct {
	store map[string]reflect.Type
}

func (h *classLoader) ClassName(i interface{}) string {
	typ := reflect.TypeOf(i)
	fullPath := ""

	typElem := typ
	if typ.Kind() == reflect.Ptr {
		fullPath = "*"
		typElem = typ.Elem()
	}
	fullPath += typElem.PkgPath() + "." + typElem.Name()

	h.store[fullPath] = typ

	return fullPath
}

func (h *classLoader) NewByType(typeName string) interface{} {
	return reflect.New(h.store[typeName]).Interface()
}

var globalClassLoader = &classLoader{
	store: make(map[string]reflect.Type, 32),
}

func ClassName(i interface{}) string {
	return globalClassLoader.ClassName(i)
}

func NewObject(typeName string) interface{} {
	return globalClassLoader.NewByType(typeName)
}
