package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//HTTPGet for http get request.
func HTTPGet(reqURL string) (resp *http.Response, err error) {
	resp, err = http.Get(reqURL)
	if err != nil {
		fmt.Printf("http.Get %v\n", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http.Status %v\n", resp.Status)
	}
	return
}

//DecodeJSONHttpResponse for json response.
func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
