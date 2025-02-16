package base

import (
	"encoding/json"
	"fmt"
	"strings"
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

func TestString(t *testing.T) {
	var builder strings.Builder
	for i := 0; i < 10; i++ {
		builder.WriteString(fmt.Sprintf("%d", i))
	}
	fmt.Println(builder.String())
}

func TestSliceBase(t *testing.T) {
	arr2 := make([]int, 0)
	var arr []int

	if arr == nil {
		fmt.Println("arr is nil")
	}
	arr2 = append(arr2, arr...)
	arr2 = append(arr2, arr...)
	arr2 = append(arr2, arr...)

	fmt.Println(len(arr2), cap(arr2))
}
