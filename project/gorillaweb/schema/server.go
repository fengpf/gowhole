package main

import (
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
	fmt.Println(person)
	b, err := json.Marshal(&Person{
		Name:  person.Name,
		Phone: person.Phone,
	})
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	fmt.Println(b)
	fmt.Fprint(w, b)
	// Do something with person.Name or person.Phone
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
