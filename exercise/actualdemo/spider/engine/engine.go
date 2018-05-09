package engine

import (
	"gowhole/exercise/actualdemo/spider/fetcher"
	"gowhole/exercise/actualdemo/spider/model"
	"log"
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
		// fmt.Printf("%s\n", body)

		parseResult := r.ParseFunc(body)
		resquests = append(resquests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("get item(%+v)", item)
		}
	}
}
