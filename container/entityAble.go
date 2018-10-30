package container

import (
	"fmt"
	"reflect"
)

type entityAble interface {
	// 	实例化实现
	buildEmptyEntity(*Container, ...interface{}) interface{}
}

// 	把任何对象转换成容器内的格式化对象
func newEntityAble(concrete interface{}) entityAble {
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

// 	构建成为新的对象
func buildEntity(able entityAble, params []interface{}) interface{} {

	if able == nil {
		panic(`nil entityAble`)
	}

	entity := able.buildEmptyEntity(nil)
	if entity == nil {
		panic(fmt.Sprint(`build entity failed for`, able))
	}

	return initWithTag(entity)
}

func initWithTag(obj interface{}) interface{} {

	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return obj
	}

	for i := 0; i < typ.NumField(); i++ {
		// log.Println("	autowired:", typ.Field(i).Name, typ.Field(i).Tag.Get("autowired"))
	}

	return obj
}
