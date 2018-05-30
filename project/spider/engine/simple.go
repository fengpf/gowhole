package engine

import (
	"log"

	"gowhole/project/spider/fetcher"
	"gowhole/project/spider/model"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...model.Request) {
	var resquests []model.Request
	for _, r := range seeds {
		resquests = append(resquests, r)
	}

	for len(resquests) > 0 {
		r := resquests[0]
		resquests = resquests[1:]
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		resquests = append(resquests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("get item(%v)", item)
		}
	}
}

func Worker(r model.Request) (model.ParseResult, error) {
	var (
		body []byte
		err  error
	)
	if body, err = fetcher.Fetch(r.URL); err != nil {
		// log.Printf("fetch url(%s) error(%v)", r.URL, err)
		return model.ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
