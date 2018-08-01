package main

import (
	"gowhole/exercise/actualdemo/simplespider/engine"
	"gowhole/exercise/actualdemo/simplespider/model"
	"gowhole/exercise/actualdemo/simplespider/parser"
)

const (
	zhenaiURL = "http://www.zhenai.com/zhenghun"
)

func main() {
	engine.Run(model.Request{
		URL:       zhenaiURL,
		ParseFunc: parser.ParseCityList,
	})
	return
}
