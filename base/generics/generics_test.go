package generics

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/duke-git/lancet/v2/convertor"
)

func TestGenericCallMethod(t *testing.T) {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

}

func TestGenericsFunc(t *testing.T) {
	m := map[string]any{
		"key": "value",
	}
	value, _ := GetSpecialValue[string](m, "key")
	fmt.Println(value)
}

var NoExists error = errors.New("not Exists")

func GetSpecialValue[T string | bool | float64 | int64](m map[string]any, key string) (T, error) {
	var t T
	val, exists := m[key]
	if !exists {
		return t, NoExists
	}
	kind := reflect.TypeOf(t).Kind()
	valElem := reflect.ValueOf(&t).Elem()
	switch kind {
	case reflect.Int64:
		toInt, _ := convertor.ToInt(val)
		valElem.SetInt(toInt)
	case reflect.String:
		toString := convertor.ToString(val)
		valElem.SetString(toString)
	case reflect.Struct:
		toString := convertor.ToString(val)
		if err := json.Unmarshal([]byte(toString), &t); err != nil {
			return t, err
		}
	default:
		return t, NoExists
	}
	return t, nil
}
