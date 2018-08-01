package main

import (
	"fmt"
	"gowhole/exercise/x/src/html/template"
	"os"
)

var (
	filename         = "a.txt"
	templateFilename = "t.txt"
	data             *string
)

func main() {
	// x := 1
	// fmt.Println(x) //prints 1
	// {
	// 	fmt.Println(x) //prints 1
	// 	x := 2
	// 	fmt.Println(x) //prints 2
	// }
	// fmt.Println(x) //prints 1 (bad if you need 2)

	// templateToFile(templateFilename, filename, &data)
	templateToFile2(templateFilename, filename, &data)
	fmt.Println(111, data)
}

func templateToFile(templateFilename string, filename string, data interface{}) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.ParseFiles(templateFilename)
	if err != nil {
		return err
	}
	return t.Execute(f, data)
}

func templateToFile2(templateFilename string, filename string, data interface{}) (err error) {
	if f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666); err == nil {
		defer f.Close()
		if t, err := template.ParseFiles(templateFilename); err == nil {
			return t.Execute(f, data)
		}
	}
	return
}
