package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"

	"gowhole/project/spider/model"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Age:       34,
		Height:    162,
		Education: "大学本科",
		Marriage:  "离异",
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	//TODO: try to start up elastic search
	//here using docker go client.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", *resp.Source)
	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if expected != actual {
		t.Errorf("got %v; expected %v\n", actual, expected)
	}
}
