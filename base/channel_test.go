package base

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelParams(t *testing.T) {
	ch := make(chan string, 5)
	go printChannel(ch)
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("%d", i)
	}
	time.Sleep(time.Second)
	close(ch)
}

func printChannel(ch chan string) {
	for a := range ch {
		fmt.Println(a)
	}
}

func TestChannelList(t *testing.T) {
	ch := make(chan string, 1000)
	addToChannel(ch)

	close(ch)

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

func addToChannel(ch chan string) {
	ch <- "hello"
	ch <- "world"
	ch <- "!"
}

func TestChannel(t *testing.T) {
	arr := getDataArr()
	l := len(arr)
	fmt.Println(l)
}

func getDataArr() []string {
	return nil
}
