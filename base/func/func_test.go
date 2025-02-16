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

func TestDeferFunc(t *testing.T) {
	type Temp struct {
		a int
		b int
	}

	temp := Temp{a: 1, b: 2}
	defer func() {
		fmt.Println("defer func")
		fmt.Println(temp)
	}()

	temp.a = 2
	temp.b = 3
	fmt.Println(temp)
}
