package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
		Addr:      "127.0.0.1:8888",
		ConfigDir: "/data/app/go/src/gowhole/exercise/gops",
	}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Minute * 5)
}
