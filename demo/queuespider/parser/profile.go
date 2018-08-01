package parser

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"gowhole/exercise/actualdemo/queuespider/model"
)

var (
	ageRe      = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe   = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	eduRe      = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	marriageRe = regexp.MustCompile(` <td><span class="label">婚况：</span>([^<]+)</td>`)
)

func ParseProfile(contents []byte, name string) model.ParseResult {
	var (
		profile           = model.Profile{}
		ageStr, heightStr string
		age, height       int
		err               error
	)
	profile.Name = name
	ageStr = extractString(contents, ageRe)
	heightStr = extractString(contents, heightRe)
	if age, err = strconv.Atoi(ageStr); err != nil {
		log.Printf("strconv.Atoi ageStr(%s)|error(%v)", ageStr, err)
	}
	profile.Age = age
	if height, err = strconv.Atoi(heightStr); err != nil {
		log.Printf("strconv.Atoi ageStr(%s)|error(%v)", ageStr, err)
	}
	profile.Height = height
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, eduRe)
	fmt.Println(profile)
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
