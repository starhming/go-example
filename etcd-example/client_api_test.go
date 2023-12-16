package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcdClient(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	defer cli.Close()

	ctx := context.Background()

	putResp, err := cli.Put(ctx, "/illusory/native", "world")
	if err != nil {
		fmt.Println(err)
	}
	marshal, _ := json.Marshal(putResp)
	fmt.Println(string(marshal))

}

func TestEtcdClientPut(t *testing.T) {
	kv, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	defer kv.Close()
	ctx := context.Background()
	putResp, err := kv.Put(ctx, "/illusory/cloud", "hello", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(putResp)
	fmt.Println(string(marshal))
}

func TestLease(t *testing.T) {
	kv, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	defer kv.Close()

	grantResp, err := kv.Grant(context.Background(), int64(time.Minute.Seconds()))
	printJsonStruct(grantResp)

	putResp, err := kv.Put(context.Background(), "/illusory/cloud/x", "ok", clientv3.WithLease(grantResp.ID))
	if err != nil {
		panic(err.Error())
	}

	printJsonStruct(putResp)
}

func TestTxn(t *testing.T) {
	kv, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	defer kv.Close()

	txn := kv.Txn(context.Background())
	commit, err := txn.If(clientv3.Compare(clientv3.Value("/illusory/cloud"), "=", "hello")).
		Then(clientv3.OpGet("/illusory/cloud")).
		Else(clientv3.OpGet("/illusory/wind", clientv3.WithPrefix())).
		Commit()

	printJsonStruct(commit)
}

func TestName(t *testing.T) {

}

func printJsonStruct(t any) {
	marshal, _ := json.Marshal(t)
	fmt.Println(string(marshal))
}
