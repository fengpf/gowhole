package persist

import (
	"context"
	"fmt"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver Got item #%d: %v", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item saver item(%v)|error(%v)", item, err)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		//must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Printf("%+v\n", resp)
	return resp.Id, nil
}
