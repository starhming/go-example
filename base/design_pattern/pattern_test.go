package design_pattern

import (
	"fmt"
	"testing"
)

type Animal interface {
	Say()
}

type Base struct{}

func (b Base) Say() {
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("汪汪汪  dog....")
}

type Bird struct{}

func (b Bird) Say() {
	fmt.Println("叽叽喳喳  bird....")
}

func TestPolymorphic(t *testing.T) {

}
