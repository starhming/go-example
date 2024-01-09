package reflect_example

import (
	"fmt"
	"reflect"
	"testing"
)

type PC interface {
	Name()
}
type Mac struct {
	name string
}

func (m Mac) Name() string {
	return m.name
}

func TestReflectType(t *testing.T) {
	var a any
	a = 1
	printTypeName(a)

	var pc any
	pc = Mac{"mac"}
	printTypeName(pc)

	printTypeName(Mac{"name"})

}

func printTypeName(a any) {
	fmt.Println(reflect.TypeOf(a).Name())
}

// 获取any的具体类型
