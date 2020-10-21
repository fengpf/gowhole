package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")
	c := cron.New()
	s1 := "*/10 * * * * *"
	c.AddFunc(s1, func() {
		log.Println("Run models.CleanAllTag...")
		//models.CleanAllTag()
	})

	s2 := "*/5 * * * * *"
	c.AddFunc(s2, func() {
		log.Println("Run models.CleanAllArticle...")
		s2 = "*/2 * * * * *"
		//models.CleanAllArticle()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
