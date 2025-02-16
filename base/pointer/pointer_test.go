package pointer

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func TestPointerDemo1(t *testing.T) {
	n := 10

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = i
	}
	fmt.Println(b)
	// [0 1 2 3 4 5 6 7 8 9]

	// 取slice的最后的一个元素
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + 9*unsafe.Sizeof(b[0]))
	// 等价于unsafe.Pointer(&b[9])
	fmt.Println(*(*int)(end))
	// 9

	runtime.Goexit()
}

func TestPointerConvert(t *testing.T) {
	txt := "TestPointerConvert"

	bytes := ([]byte)(txt)

	fmt.Println(bytes)
	bytes[0] = 80
	fmt.Println(bytes)

	fmt.Println(string(bytes))
}

func byteArrConvertString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}
