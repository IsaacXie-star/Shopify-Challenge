package dao

import (
	"Shopify-Challenge/db"
	"log"
	"sync"
)

// ItemListDao is a Data Access Object
type ItemListDao struct {
}

var itemListDao *ItemListDao
var itemListDaoCreateOnce sync.Once

func NewItemListDao() *ItemListDao {
	itemListDaoCreateOnce.Do(
		func() {
			itemListDao = &ItemListDao{}
		})
	return itemListDao
}

//QueryShopItemList queries items from the database base on the conditions
func (dao *ItemListDao) QueryShopItemList(cods map[string]string) ([]*db.Item, int64, error) {
	equalCods := make(map[string]string)
	greaterCods := make(map[string]string)
	lessCods := make(map[string]string)
	if id, ok := cods["id"]; ok {
		equalCods["id"] = id
	}
	if name, ok := cods["name"]; ok {
		equalCods["name"] = name
	}
	if category, ok := cods["category"]; ok {
		equalCods["category"] = category
	}
	if status, ok := cods["status"]; ok {
		equalCods["status"] = status
	}
	if maxPrice, ok := cods["max_price"]; ok {
		greaterCods["max_price"] = maxPrice
	}
	if minPrice, ok := cods["min_price"]; ok {
		lessCods["min_price"] = minPrice
	}

	shopItemList, total, err := db.SelectWhereItem(equalCods, greaterCods, lessCods)
	if err != nil {
		log.Printf("[QueryShopItemList] db.SelectWhereItem error, err=%v", err)
		return nil, 0, err
	}
	return shopItemList, total, nil
}
