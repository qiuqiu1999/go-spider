package persist

import (
	"fmt"
	"go-spider/engine"
)

func ItemSaver() chan engine.Item {
	itemChan := make(chan engine.Item)

	go func() {
		for  {
			item := <- itemChan
			fmt.Printf("Got item %+v \n", item)
		}
	}()

	return itemChan
}
