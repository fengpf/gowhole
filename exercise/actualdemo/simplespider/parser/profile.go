package parser

import (
	"regexp"
	"strconv"

	"gowhole/exercise/actualdemo/simplespider/model"
)

var (
	ageRe      = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe   = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	eduRe      = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	marriageRe = regexp.MustCompile(` <td><span class="label">婚况：</span>([^<]+)</td>`)
)

func ParseProfile(contents []byte, name string) model.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, eduRe)
	result := model.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
