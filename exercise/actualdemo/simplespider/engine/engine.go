package engine

import (
	"log"

	"gowhole/exercise/actualdemo/simplespider/fetcher"
	"gowhole/exercise/actualdemo/simplespider/model"
)

func Run(seeds ...model.Request) {
	var resquests []model.Request
	for _, r := range seeds {
		resquests = append(resquests, r)
	}

	for len(resquests) > 0 {
		r := resquests[0]
		resquests = resquests[1:]

		var (
			body []byte
			err  error
		)
		log.Printf("fetching %s\n", r.URL)
		if body, err = fetcher.Fetch(r.URL); err != nil {
			log.Printf("fetch url(%s) error(%v)", r.URL, err)
		}
		parseResult := r.ParseFunc(body)
		resquests = append(resquests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("get item(%+v)", item)
		}
	}
}
