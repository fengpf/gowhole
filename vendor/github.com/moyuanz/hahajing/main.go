package main

import (
	"github.com/moyuanz/hahajing/com"
	"github.com/moyuanz/hahajing/door"
	"github.com/moyuanz/hahajing/kad"
	"github.com/moyuanz/hahajing/web"
)

var kadInstance kad.Kad
var webInstance web.Web
var doorInstance door.Door
var keywordManager = com.NewKeywordManager()

func main() {
	kadInstance.Start()
	doorInstance.Start(keywordManager)

	webInstance.Start(kadInstance.SearchReqCh, doorInstance.KeywordCheckReqCh, keywordManager)
}
