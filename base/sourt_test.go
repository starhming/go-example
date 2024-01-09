package base

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

func TestSortFunc(t *testing.T) {
	type A struct {
		a int `json:"a"`
	}

	arr := []A{
		{1},
		{5},
		{3},
		{4},
		{2},
		{0},
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].a < arr[j].a
	})

	marshal, _ := json.Marshal(arr)
	fmt.Println(string(marshal))

}
