package schema

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/gorilla/schema"
)

type Person struct {
	Name  string
	Phone string
}

func myHTTPRequest() {
	var encoder = schema.NewEncoder()
	person := Person{"Jane Doe", "555-5555"}
	form := url.Values{}
	err := encoder.Encode(person, form)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	// Use form values, for example, with an http client
	client := new(http.Client)
	res, err := client.PostForm("http://localhost:8000", form)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func Test_encode(t *testing.T) {
	myHTTPRequest()
}

func Test_decode(t *testing.T) {
	r := http.NewServeMux()
	r.HandleFunc("/", myHandler)
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
	// Do something with person.Name or person.Phone
}
