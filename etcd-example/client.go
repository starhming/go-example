package main

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func CreateEtcdClient() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	return cli
}
