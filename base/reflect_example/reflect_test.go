package reflect_example

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/tidwall/gjson"
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

type Employee struct {
	EmployeeId string `json:"employee_id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

// 设置值
func TestInvokeByName(t *testing.T) {
	//e := &Employee{"1", "Mike", 30}
	data := `{"employee_id": "111", "name": 22, "age": "12.8"}`

	e := &Employee{}
	val := reflect.ValueOf(e).Elem()
	typ := reflect.TypeOf(e).Elem()

	for i := 0; i < val.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("json")
		v := gjson.Get(data, tag).Value()

		kind := typ.Field(i).Type.Kind()
		switch kind {
		case reflect.String:
			vStr := convertor.ToString(v)
			val.Field(i).SetString(vStr)
		case reflect.Int:
			vInt, err := convertor.ToInt(v)
			if err != nil {
				fmt.Printf("Field: %s, type is invalid: %s /n", tag, err.Error())
			}
			val.Field(i).SetInt(vInt)
		}
	}

	marshal, _ := json.Marshal(e)
	fmt.Println(string(marshal))
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSetValueToStruct(t *testing.T) {
	p := &Person{}
	v := reflect.ValueOf(p).Elem()

	name := "hm"
	age := 24

	fmt.Println(v.NumField())

	//v.FieldByName("Name").Set(reflect.ValueOf(name))
	//v.FieldByName("Name").SetString(name)
	v.Field(0).SetString(name)

	v.FieldByName("Age").Set(reflect.ValueOf(age))

	fmt.Println(*p)

}
