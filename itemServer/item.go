package itemServer

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

func saveItem(item interface{}) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.64.130:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.Index().Index("zhenai").Type("info").BodyJson(item).Do(context.Background())

	fmt.Printf("%+v\n", resp)
}
