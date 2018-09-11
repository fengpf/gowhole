package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var ErrTimeout = errors.New("received timeout")     //任务执行超时返回
var ErrInterrupt = errors.New("received interrupt") //任务接受到中断返回

//Runner 在给定的一段超时时间内执行一组任务,并且在操作系统发送中断信号的时候结束这些任务
type Runner struct {
	interrupt chan os.Signal   //结束信号
	complete  chan error       //通知任务完成
	timeout   <-chan time.Time //通知任务超时
	tasks     []func(int)      //任务列表
}

//New 返回一个准备使用的runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add 添加任务
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//Start 执行任务
func (r *Runner) Start() error {

	//期望接收的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	//使用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	//监听处理完成的信号和超时信号
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

//run 执行每一个已经注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

//gotInterrupt 验证是否收到中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	//当中断信号被触发的时候发出信号
	case <-r.interrupt:
		//停止后续接收的任何信号
		signal.Stop(r.interrupt)
		return true
	default: //防止阻塞
		return false
	}
}
