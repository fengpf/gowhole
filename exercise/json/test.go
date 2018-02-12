package main

import (
	"encoding/json"
)

//SucExpctionRespItem ...
type SucExpctionRespItem struct {
	Context string      `json:"context"`
	Body    interface{} `json:"body"`
	Status  int         `json:"status"`
}

//SucExpctionResp ...
type SucExpctionResp struct {
	Result SucExpctionRespItem `json:"result"`
}

//Result ..
type Result struct {
	Exception     string      `json:"exception"`
	Code          int         `json:"code"`
	OfficeResult  interface{} `json:"OfficeResult"`
	Context       string      `json:"context"`
	OfficeMessage string      `json:"OfficeMessage"`
	Error         string      `json:"error"`
	Message       string      `json:"message"`
	Status        int         `json:"status"`
}

func main() {
	var result Result
	res := `{
		"exception":"aaaaaaaaa",
		"code":666666,
		"OfficeResult":{"total":"1","success":false,"paramList":[],"messageCode":"","message":"账户不存在"},
		"context":"success",
		"OfficeMessage":"账户不存在",
		"error":"INTERNAL_SERVER_ERROR",
		"message":"账户不存在",
		"status":500
	}`
	err := json.Unmarshal([]byte(res), &result)
	// fmt.Printf("%v\n", err)
	if err != nil {
		println("aaa")
	}
	println(result.OfficeMessage)
}
