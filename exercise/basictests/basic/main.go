package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int(), (rand.Int() % 1000), (rand.Int()%1000) < 100)
}
