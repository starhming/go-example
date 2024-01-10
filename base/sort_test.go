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

type Timer interface {
	GetTime() int64
}

func TimerSortAsc[T Timer](arr []T) []T {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].GetTime() <= arr[j].GetTime()
	})
	return arr
}

type Feature struct {
	eventTime int64
}

func (f Feature) GetTime() int64 {
	return f.eventTime
}

func TestCustomerTimeSort(t *testing.T) {
	arr := []Feature{
		{1},
		{8},
		{3},
		{5},
		{0},
	}

	asc := TimerSortAsc(arr)

	fmt.Println(asc)
}
