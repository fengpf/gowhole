package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertTpBin(n int) string {
	res := ""
	if n == 0 {
		res = "0"
		return res
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertTpBin(5),  //101
		convertTpBin(13), //1011
		convertTpBin(343434),
		convertTpBin(0),
	)

	printFile("abc.txt")
	forever()
}
