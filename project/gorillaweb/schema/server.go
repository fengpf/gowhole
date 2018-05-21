package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/schema"
)

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	var decoder = schema.NewDecoder()
	err := r.ParseForm()
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	var person Person
	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&person, r.PostForm)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	p := &Person{
		Name:  person.Name,
		Phone: person.Phone,
	}
	fmt.Printf("server get data(%v)\n", p)
	b, err := json.Marshal(p)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	var (
		buffer bytes.Buffer
		buff   bytes.Buffer
	)
	buff.Write(b)
	if err := json.Compact(&buffer, buff.Bytes()); err != nil {
		println(err) //error message: invalid character ',' looking for beginning of value
		return

	}
	w.Write(buffer.Bytes())
	// Do something with person.Name or person.Phone
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
