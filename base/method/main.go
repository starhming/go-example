package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect) growWidth() {
	r.width = r.width * 2
}

func (r *rect) growHeight() {
	r.height = r.height * 2
}

func main() {
	r := rect{width: 10, height: 5}
	rp := &r

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	r.growHeight()
	r.growWidth()
	fmt.Println("r width: ", r.width, " height: ", r.height)

	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	rp.growWidth()
	rp.growHeight()
	fmt.Println("rp width: ", rp.width, " height: ", rp.height)
}
