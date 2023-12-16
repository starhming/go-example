package container

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()
	l.PushFront(4)
}
