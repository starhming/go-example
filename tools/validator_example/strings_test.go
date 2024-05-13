package validator_example

import (
	"fmt"
	"testing"
)

func TestStringValidate(t *testing.T) {
	if err := validate.Var("stream-1111", "startswith=stream"); err != nil {
		fmt.Println(err.Error())
	}

	if err := validate.Var(" stream-1111", "startswith=stream"); err != nil {
		fmt.Println(err.Error())
	}

}
