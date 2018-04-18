package main

import (
	"encoding/json"
	"fmt"
	"gowhole/syslearning/queue/schedule"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	MaxWorker       = 2
	MaxQueue        = 2
	MaxLength int64 = 2048
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Read the body into a string for json decoding
	var content = &schedule.PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body), content)
	if err != nil {
		err = fmt.Errorf("an error occured while deserializing message")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(content.Payloads)
	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {
		// let's create a job with the payload
		work := schedule.Job{Payload: payload}
		fmt.Println("sending payload  to workque")
		// Push the work onto the queue.
		schedule.JobQueue <- work
		fmt.Println("sent payload  to workque")
	}
	w.WriteHeader(http.StatusOK)
}
func main() {
	schedule.JobQueue = make(chan schedule.Job, MaxQueue)

	dp := schedule.NewDispatcher(MaxWorker)
	dp.Run()

	http.HandleFunc("/payload", payloadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("starting listening for payload messages")
	} else {
		err = fmt.Errorf("an error occured while starting payload server %s", err.Error())
	}
	fmt.Println(err)
	// time.Sleep(time.Minute)
}
