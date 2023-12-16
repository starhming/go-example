package main

import "fmt"

type Resource struct {
	Count  int64
	Core   int64
	Memory int64
}

func main() {
	//arr := []Resource{
	//	{20, 10, 14336},
	//	{200, 10, 10240},
	//	{62, 7, 17408},
	//	{12, 4, 10240},
	//}
	//
	//var core, memory int64
	//for _, r := range arr {
	//	c := r.Core * r.Count
	//	m := r.Memory * r.Count
	//
	//	core += c
	//	memory += m
	//}
	//
	//fmt.Println("core: ", core)
	//fmt.Println("memory: ", memory/1024) // T

	cursor := 0
	for i := cursor; i < 10; i++ {
		fmt.Printf("i: %d, cursor: %d\r\n", i, cursor)
		cursor++
	}
	fmt.Println(cursor)
}
