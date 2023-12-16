package generics

import (
	"fmt"
	"testing"
)

func TestGenericCallMethod(t *testing.T) {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
