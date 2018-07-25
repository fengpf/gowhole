package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func catch(e error) {
	if e != nil {
		log.Fatalln(e)
	}
} // END catch

func main() {
	// Cookie jar
	cookies, err := cookiejar.New(nil)
	catch(err)

	// Custom client
	client := &http.Client{
		Jar:     cookies,
		Timeout: time.Second * 60,
	}

	// Data to be send
	data := url.Values{}
	// Fake values
	data.Add("user", "john")
	data.Add("passwd", "123")

	// Data to buffer
	buffData := bytes.NewBuffer([]byte(data.Encode()))

	// Fake URL
	strURL := "http://www.google.com/"

	// Generating request
	req, err := http.NewRequest("POST", strURL, buffData)
	catch(err)

	// Setting required headers
	req.Header.Set("Postman-Token", "bc6825dd-d6d5-2df8-204d-a70824a74d4b")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive") // Force Connection: keep-alive

	//req.Close = false // Another try to get Connection: keep-alive

	// Making the request and getting response
	res, err := client.Do(req)
	catch(err)

	// Reading response body
	b := new(bytes.Buffer)
	b.ReadFrom(res.Body)
	strBody := b.String()

	// Fake string
	strSuccess := "logged in"

	isLoggedIn := strings.Contains(strBody, strSuccess)

	fmt.Printf("Is logged in? %v\n", isLoggedIn)
} // END main
