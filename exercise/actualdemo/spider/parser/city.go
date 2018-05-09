package parser

import (
	"regexp"

	"gowhole/exercise/actualdemo/spider/model"
)

const (
	cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
)

func ParseCity(contents []byte) model.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := model.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests,
			model.Request{
				URL: string(m[1]),
				ParseFunc: func(c []byte) model.ParseResult {
					return ParseProfile(c, name)
				},
			})
	}
	return result
}
