package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

var ErrPoolClosed = errors.New("pool has been closed")

//Pool 管理一组可以安全地在多个goroutine间共享的资源，被管理资源必须实现 io.Closer接口
type Pool struct {
	m        sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

//创建一个管理资源的池子，这个池子需要传入可以分配资源的函数，并规定池子的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{
		factory:  fn,
		resource: make(chan io.Closer, size),
	}, nil
}

//Acquire 从资源池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	//检查是否有空闲的资源
	case r, ok := <-p.resource:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
		//如果没有空闲的资源，在分配新资源
	default:
		log.Println("Acquire:", " New Resource")
		return p.factory()
	}
}

//Release 回收已使用的资源到资源池
func (p *Pool) Release(r io.Closer) {
	//保证回收和close操作的安全
	p.m.Lock()

	//如果资源池已经被关闭，则销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	//回收资源到队列
	case p.resource <- r:
		log.Println("Release:", "In Queue")

		//如果队列已满则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}

	p.m.Unlock()
}

//Close 停止资源池的工作，并且关闭所有资源
func (p *Pool) Close() {
	//保证线程安全操作
	p.m.Lock()

	//如果资源池已被关闭则啥也不做
	if p.closed {
		return
	}

	//关闭资源池
	p.closed = true

	//在清空资源池通道之前关闭此通道,否则发生死锁
	close(p.resource)

	//关闭资源
	for r := range p.resource {
		r.Close()
	}

	p.m.Unlock()
}
