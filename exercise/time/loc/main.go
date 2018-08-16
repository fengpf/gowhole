package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC822, "13 Aug 18 11:36 CST")
	if err != nil {
		log.Printf("time.Parse error(%v)", err)
		return
	}

	loc, err := time.LoadLocation("UTC")

	// loc, err := time.LoadLocation(time.Local.String())

	if err != nil {
		log.Printf("time.LoadLocation error(%v)", err)
		return
	}
	fmt.Println(1111, t)

	fmt.Println(2222, loc)

	fmt.Println(rand.Intn(1000))
}
