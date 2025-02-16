// hello.go
package main

/*
#include <stdio.h>
#include <stdlib.h>

static void SayHello(const char* s) {
    puts(s);
}

void myprint(char* s) {
    printf("%s\n", s);
}
*/
import "C"
import "unsafe"

func main() {
	C.SayHello(C.CString("Hello, World\n"))

	cs := C.CString("Hello, World\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
