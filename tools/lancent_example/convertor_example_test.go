package lancent_example

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/convertor"
)

func TestConvertor(t *testing.T) {
	result1, err := convertor.ToFloat("")
	if err != nil {
		fmt.Printf("/t ToFloat: '', %s", err.Error())
	}
	result2, err := convertor.ToFloat("abc")
	if err != nil {
		fmt.Errorf("ToFloat: '', %s", err.Error())
	}
	result3, _ := convertor.ToFloat("-1")

	result4, _ := convertor.ToFloat("-.11")
	result5, _ := convertor.ToFloat("1.23e3")
	result6, _ := convertor.ToFloat(true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// 0
	// 0
	// -1
	// -0.11
	// 1230
	// 0
}
