package base

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const (
	a = -1
	b = iota
	c
	d
)

func TestBaseIota(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func TestMapExample(t *testing.T) {
	m := make(map[string]map[string]int)

	if _, exists := m["a"]; !exists {
		m["a"] = make(map[string]int)
	}
	for i := 0; i < 10; i++ {
		m["a"][fmt.Sprintf("%d", i)] = i
	}

	marshal, _ := json.Marshal(m)
	fmt.Println(string(marshal))
}

type DiagData map[int]string

func TestStructAlias(t *testing.T) {
	m := map[int]string{1: "1", 2: "2"}

	var dd DiagData
	dd = m

	fmt.Println(dd)
}

func TestName(t *testing.T) {
	var aa []int
	arrAdd(aa)
	fmt.Println(aa)
}

func arrAdd(arr []int) {
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
}

func TestFieldType(t *testing.T) {
	type PC struct {
		Cpu  any    `json:"cpu,omitempty"`
		Name string `json:"name"`
	}

	pc := PC{Cpu: 100.1, Name: "Alienware"}

	//typ := reflect.TypeOf(pc)
	val := reflect.ValueOf(pc)

	v := val.Field(0).Interface()
	fmt.Println(v)

	v = val.Field(1).Interface()
	fmt.Println(v)

}
