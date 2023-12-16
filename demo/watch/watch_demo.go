package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	addr     = ":8080"
	notifyCh = make(chan int)
)

func Production() {
	for i := 0; i < 5; i++ {
		notifyCh <- i + 1
		<-time.Tick(time.Second * 3)
	}
	close(notifyCh)
}

func Watch(w http.ResponseWriter, r *http.Request) {
	flusher := w.(http.Flusher)
	for {
		v, ok := <-notifyCh
		if !ok {
			flusher.Flush()
			return
		}
		fmt.Fprintf(w, "%v\n", v)
		flusher.Flush()
	}
}

func main() {
	go Production()

	http.HandleFunc("/watch", Watch)
	fmt.Printf("listem addr: %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
