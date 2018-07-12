package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
	endpoints      = []string{"127.0.0.1:2379"}
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		println(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, "/test/hello", "world")
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "/test/hello")
	cancel()

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	_, err = cli.Put(context.TODO(), "key", "xyz")
	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Txn(ctx).
		If(clientv3.Compare(clientv3.Value("key"), ">", "abc")).
		Then(clientv3.OpPut("key", "XYZ")).
		Else(clientv3.OpPut("key", "ABC")).
		Commit()
	cancel()

	rch := cli.Watch(context.Background(), "/test/hello", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	if err != nil {
		println(err)
	}

}
