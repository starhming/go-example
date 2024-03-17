package mapstruct_example

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  any
}

func TestMapStructExampleDemo(t *testing.T) {
	m := map[string]any{"Name": "aa", "Age": "100"}

	p := Person{}
	err := mapstructure.Decode(m, &p)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(p)
}
