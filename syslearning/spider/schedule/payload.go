package schedule

import (
	"fmt"
	"net/http"
)

type PayloadCollection struct {
	Payloads []Payload `json:"data"`
}

type Payload struct {
	// [redacted]
	A int
}

func (p *Payload) UploadToS3() {
	fmt.Println(6666666)
	return
}

func PayloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Read the body into a string for json decoding
	payloads := make([]Payload, 0, 2)
	payload := Payload{A: 1}
	payloads = append(payloads, payload)
	// payload2 := Payload{A: 2}
	// payloads = append(payloads, payload2)
	var content = &PayloadCollection{Payloads: payloads}
	// err := json.NewDecoder(io.LimitReader(r.Body, 10000)).Decode(&content)
	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {
		println(payload.A)
		// let's create a job with the payload
		work := Job{Payload: payload}

		// Push the work onto the queue.
		JobQueue <- work
	}
	w.WriteHeader(http.StatusOK)
}
