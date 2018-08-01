package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"goutil/calendar/cron"
)

var (
	Conf configuration
	err  error
)

type configuration struct {
	Run string
}

func loadConfig() (conf configuration) {
	Conf, err = readConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}
	log.Printf("start app...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Printf("app exit...")
			return
		case syscall.SIGHUP:
			Conf, err = readConfig()
			if err != nil {
				fmt.Println("Error:", err)
			}
			log.Printf("reload Config(%s)\n", Conf.Run)
			return
		default:
			return
		}
	}
}

func readConfig() (conf configuration, err error) {
	file, _ := os.Open("./job.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf = configuration{}
	err = decoder.Decode(&conf)
	return conf, err
}

func main() {
	// callBack(2, Add)

	i := 0
	c := cron.New()
	spec := loadConfig().Run
	log.Printf("spec(%s)\n", spec)
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running : ", i)
	})
	c.AddJob(spec, TestJob{})
	// c.AddJob(spec, Test2Job{})
	c.Start()
	defer c.Stop()
	select {}
}

type TestJob struct {
}

// type Test2Job struct {
// }

func (this TestJob) Run() {
	fmt.Println("testJob1...")
}

// func (this Test2Job) Run() {
// 	fmt.Println("testJob2...")
// }

func Add(a, b int) {
	fmt.Printf("a(%d),b(%d),a+b(%d)\n", a, b, a+b)
}

// type myFunc func(int, func(int, int)) interface{}

func callBack(a int, f func(i, j int)) {
	f(a, 2)
}
