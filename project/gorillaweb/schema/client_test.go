package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/gorilla/schema"
)

func Test_encode(t *testing.T) {
	var encoder = schema.NewEncoder()
	person := Person{"JaneDoe", "555-5555"}
	form := url.Values{}
	err := encoder.Encode(person, form)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	// Use form values, for example, with an http client
	client := new(http.Client)
	res, err := client.PostForm("http://localhost:8080", form)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var p Person
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Printf("json.Unmarshal error(%v)\n", err)
		return
	}
	fmt.Println(string(body), p.Name, p.Phone)
}
