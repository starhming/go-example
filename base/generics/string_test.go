package generics

import (
	"fmt"
	"testing"
)

func TestStringLogic(t *testing.T) {
	a := "10"
	fmt.Println(a > "100")
	fmt.Println(a > "5")
	fmt.Println(a < "10")

	fmt.Println(1 > 2 && false)
}
