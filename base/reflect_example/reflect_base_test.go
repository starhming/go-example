package reflect_example

import (
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {

}

func TestSetValue(t *testing.T) {
	a := 1

	v := reflect.ValueOf(a)

	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	v = reflect.ValueOf(&a)
	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	v = v.Elem() // element value
	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	// set
	v.SetInt(2)
	fmt.Println("after set, v:", v)

	newValue := reflect.ValueOf(3)
	v.Set(newValue)
	fmt.Println("after set, v:", v)

}
