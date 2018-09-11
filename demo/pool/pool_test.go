package pool

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxGoroutines   = 10 //使用的goroutine数量
	pooledResources = 3  //池中资源数量
)

var idCounter int32 //用来给每个连接分配独一无二的id

//dbConn 模拟共享的资源
type dbConn struct {
	ID int32
}

func (db *dbConn) Close() error {
	log.Println("Close: Connection", db.ID)
	return nil
}

//createConn 是一个工厂函数，当需要一个连接的时候则资源池会调用这个函数进行连接创建
func createConn() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConn{id}, nil
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	//创建连接池
	p, err := New(createConn, pooledResources)
	if err != nil {
		log.Fatalln(err)
	}

	//使用池子里面的连接完成查询
	for query := 0; query < maxGoroutines; query++ {
		//每个goroutine需要复制一份自己要查询的副本，不然所有的查询都会共享
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()

	//关闭连接池
	log.Println("Close pool")
	p.Close()
}

//performQueries  用来测试连接的资源池
func performQueries(query int, p *Pool) {
	//从池子请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	//用等待模拟查询响应
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	log.Printf("QueryID(%d) ConnID(%d)\n", query, conn.(*dbConn).ID)

	//将该连接放回池子
	p.Release(conn)
}
