package base

import (
	"encoding/json"
	"fmt"
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
