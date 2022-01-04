package persist

import (
	"fmt"
	"go-spider/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	itemChan := make(chan engine.Item)

	go func() {
		for  {
			item := <- itemChan
			fmt.Printf("Got item %+v \n", item)
		}
	}()

	return itemChan, nil
}
