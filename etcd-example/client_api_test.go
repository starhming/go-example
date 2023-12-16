package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/starshm/go-example/util"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcdClient(t *testing.T) {
	cli := CreateEtcdClient()
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
	cli := CreateEtcdClient()
	defer cli.Close()
	ctx := context.Background()
	putResp, err := cli.Put(ctx, "/illusory/cloud", "hello", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(putResp)
	fmt.Println(string(marshal))
}

func TestLease(t *testing.T) {
	cli := CreateEtcdClient()
	defer cli.Close()

	grantResp, err := cli.Grant(context.Background(), int64(time.Minute.Seconds()))
	util.PrintJsonStruct(grantResp)

	// KeepAlive:自动定时的续约某个租约。
	// KeepAliveOnce:为某个租约续约一次
	// cli.KeepAlive()
	// cli.KeepAliveOnce()

	putResp, err := cli.Put(context.Background(), "/illusory/cloud/x", "ok", clientv3.WithLease(grantResp.ID))
	if err != nil {
		panic(err.Error())
	}

	util.PrintJsonStruct(putResp)
}

func TestTxn(t *testing.T) {
	cli := CreateEtcdClient()
	defer cli.Close()

	txn := cli.Txn(context.Background())
	commit, err := txn.If(clientv3.Compare(clientv3.Value("/illusory/cloud"), "=", "hello")).
		Then(clientv3.OpGet("/illusory/cloud")).
		Else(clientv3.OpGet("/illusory/wind", clientv3.WithPrefix())).
		Commit()
	if err != nil {
		fmt.Printf("txn failed, err: %s", err.Error())
	}
	util.PrintJsonStruct(commit)
}

func TestWatch(t *testing.T) {
	cli := CreateEtcdClient()
	defer cli.Close()
	watchChan := cli.Watch(context.Background(), "/hello")
	for wr := range watchChan {
		for _, e := range wr.Events {
			switch e.Type {
			case clientv3.EventTypePut:
				fmt.Printf("watch event put-current: %#v \n", string(e.Kv.Value))
			case clientv3.EventTypeDelete:
				fmt.Printf("watch event delete-current: %#v \n", string(e.Kv.Value))
			default:
			}
		}
	}
}
