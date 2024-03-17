package main

import (
	"fmt"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var sg sync.WaitGroup

type Counter struct {
	count int
}

func (m *Counter) Incr() {
	m.count++
}

func (m *Counter) Count() int {
	return m.count
}

func main() {
	endpoints := []string{"http://127.0.0.1:32379"}
	// 初始化etcd客户端
	client, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	counter := &Counter{}

	sg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			// 这里会生成租约，默认是60秒
			session, err := concurrency.NewSession(client)
			if err != nil {
				panic(err)
			}
			defer session.Close()

			locker := concurrency.NewLocker(session, "/my-test-lock")
			locker.Lock()
			counter.Incr()
			locker.Unlock()
			sg.Done()
		}()
	}
	sg.Wait()

	fmt.Println("count:", counter.Count())
}
