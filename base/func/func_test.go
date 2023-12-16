package _func

import (
	"fmt"
	"testing"
	"time"
)

func TestFuncCallBack(t *testing.T) {
	a := "a"
	b := "b"
	go process("xxxx", func() {
		time.Sleep(time.Second)
		fmt.Println(a)
		fmt.Println(b)
	})
	time.Sleep(time.Second * 5)
}

func process(taskId string, callback func()) {
	fmt.Println(taskId)
	callback()
}
