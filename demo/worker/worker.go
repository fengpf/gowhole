package worker

import "sync"

//Worker 必须实现接口才能使用工作池
type Worker interface {
	Task()
}

//Pool 提供一个goroutine 池子，这个池子可以完成任何一个已提交的任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

//New 创建goroutine池子
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

//Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//Close 等待所有goroutine 停止工作
func (p *Pool) Close() {
	close(p.work)
	p.wg.Wait()
}
