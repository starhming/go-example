package base

import (
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
