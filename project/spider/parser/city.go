package parser

import (
	"regexp"

	"gowhole/project/spider/model"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURLRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) model.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := model.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests,
			model.Request{
				URL: string(m[1]),
				ParseFunc: func(c []byte) model.ParseResult {
					return ParseProfile(c, name)
				},
			})
	}
	matches = cityURLRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, model.Request{
			URL:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
