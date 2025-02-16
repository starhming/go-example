package main

/*
#include <stdio.h>
extern void ACFunction();
*/
import "C"
import "fmt"

//export AGoFunction
func AGoFunction() {
	fmt.Println("AGoFunction()")
}

func Example() {
	C.ACFunction()
}

func main() {
	C.ACFunction()
}
