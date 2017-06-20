package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

var (
	_zeroWidthReg = regexp.MustCompile(`[\x{200b}]+`)
	_nocharReg    = []*regexp.Regexp{
		regexp.MustCompile(`[\p{Hangul}]+`),  // kr
		regexp.MustCompile(`[\p{Tibetan}]+`), // tibe
		regexp.MustCompile(`[\p{Arabic}]+`),  // arabic
	}
)

func main() {
	defer timeCost(time.Now())
	// var err error
	// txt := strconv.FormatInt(222222, 10)
	// wm := watermark.NewWatermark("./img/uid_mark.png", txt, 32)
	// if wm.Draw(false) != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return
	// }

	// txt := "golang水印大多数"
	// wm := watermark.NewWatermark("./img/mark.png", txt, 32)
	// if wm.Draw(true) != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return
	// }
	//imgprocessing.HTTPPrint()
	//fmt.Printf("%s\n", "create watermark success ...")
	//rpcx.Start()
	//signalHandler()
	ss := "   "
	checkTitle(ss)
}

func timeCost(start time.Time) {
	terminal := time.Since(start)
	fmt.Println(terminal)
}

func signalHandler() {
	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		println("get a signal")
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			println("exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}

func checkTitle(title string) (ct string, ok bool) {
	title = strings.TrimSpace(title)
	println(utf8.RuneCountInString(title))
	if utf8.RuneCountInString(title) > 80 {
		return
	}
	for _, reg := range _nocharReg {
		if reg.MatchString(title) {
			return
		}
	}
	ct = _zeroWidthReg.ReplaceAllString(title, "")
	ok = true
	return
}
