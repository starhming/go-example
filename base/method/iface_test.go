package main

import (
	"fmt"
	"testing"
	"time"
)

type Animal struct {
	Age int
}

func (a *Animal) Move() {
	fmt.Println("animal move")
}

type Dog struct {
	Animal
}

func (d *Dog) Move() {
	fmt.Println("dog move")
}

func TestMethodCall(t *testing.T) {
	dog := Dog{}
	dog.Move()

	dog.Animal.Move()
}

func TestDefer(t *testing.T) {
	doWork()
}

func doWork() {
	now := time.Now()
	a := Animal{100}
	defer printWork(now, a)
	time.Sleep(time.Second * 10)
	a.Age = 100
}

func printWork(lastTime time.Time, a Animal) {
	fmt.Println("now: ", time.Now().UnixMilli())
	fmt.Println("last time: ", lastTime.UnixMilli())
	fmt.Println(a)
}
