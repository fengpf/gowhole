package parser

import (
	"io/ioutil"
	"testing"
)

const (
	resSize = 470
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("content.html")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", contents)
	res := ParseCityList(contents)

	expectedURLS := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	for i, u := range expectedURLS {
		if res.Requests[i].URL != u {
			t.Errorf("expect url #%d:  %s"+"requests but had %s", i, u, res.Requests[i].URL)
		}
	}
	for i, city := range expectedCities {
		if res.Items[i].(string) != city {
			t.Errorf("expect url #%d:  %s"+"requests but had %s", i, city, res.Items[i].(string))
		}
	}
	if len(res.Requests) != resSize {
		t.Errorf("res should have %d"+"requests but had %d", resSize, len(res.Requests))
	}
	if len(res.Items) != resSize {
		t.Errorf("res should have %d"+"requests but had %d", resSize, len(res.Items))
	}
}
