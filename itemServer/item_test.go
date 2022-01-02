package itemServer

import (
	"go-spider/model"
	"testing"
)

func Test_saveItem(t *testing.T) {
	item := model.Profile{
		Name:      "风和叶子",
		Gender:    "女士",
		Age:       70,
		Height:    150,
		Income:    "重庆",
		Marriage:  "丧偶",
		Education: "高中及以下",
	}

	saveItem(item)
}
