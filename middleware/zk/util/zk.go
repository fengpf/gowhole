package util

import (
	"fmt"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	timeOut = 20
)

var hosts []string = []string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"} // the zk server list

func GetConnect() (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect(hosts, timeOut*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	var (
		stat *zk.Stat
		wg   sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		list, stat, err = conn.Children("/go_servers")
		if err != nil {
			panic(err)
		}
		println(11111111, list, stat)
	}()
	wg.Wait()
	return
}
