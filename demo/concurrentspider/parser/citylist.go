package parser

import (
	"regexp"

	"gowhole/exercise/actualdemo/concurrentspider/model"
)

const (
	cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
)

func ParseCityList(contents []byte) model.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := model.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests,
			model.Request{
				URL:       string(m[1]),
				ParseFunc: ParseCity,
			})
	}
	return result
}
