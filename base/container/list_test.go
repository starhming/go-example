package container

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()
	l.PushFront(4)
	l.PushBack(5)
	l.PushBack(7)
	l.PushBack(8)
	l.PushFront(1)

	sliceAny := listToSlice(l)
	arr := make([]int, l.Len(), l.Len())
	for i, v := range sliceAny {
		arr[i] = v.(int)
	}
	fmt.Println(arr)

}

func listToSlice(l *list.List) []interface{} {
	// 创建一个 slice
	s := make([]interface{}, l.Len())

	// 遍历 list，将元素添加到 slice 中
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		s[i] = e.Value
		i++
	}

	return s
}
