package persist

import (
	"fmt"
	"go-spider/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	itemChan := make(chan engine.Item)

	go func() {
		countItem := 0
		for {
			item := <-itemChan
			countItem++
			fmt.Printf("Got item id:%d %+v \n", countItem, item)
		}
	}()

	return itemChan, nil
}
