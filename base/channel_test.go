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
