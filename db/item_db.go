package db

import (
	"errors"
	"log"
	"strconv"
	"sync/atomic"
)

// In a typical application, this would instead be a database
// using MySQL, Oracle etc.
// In this application, we use memory storage to simulate a database

type Item struct {
	Id              int64   `json:"id"`               // unique id of an item
	Name            string  `json:"name"`             // name of the item
	Category        string  `json:"category"`         // store item's category
	Price           float64 `json:"price"`            // price of the item
	Status          string  `json:"status"`           // "0": item unavailable, deleted "1": item available
	DeletionComment string  `json:"deletion_comment"` // DeletionComment of why item is deleted (unavailable)
	Description     string  `json:"description"`      // brief description of the item
}

var shopItemList []*Item
var itemId int64

func init() {
	log.Println("Start connection to Database")
	shopItemList = make([]*Item, 0)
	itemId = 0
	addPredefinedItem()
	log.Println("Database connection established!")
}

func getId() int64 {
	// id needs to be distinct, use atomic add to avoid data races
	atomic.AddInt64(&itemId, 1)
	return itemId
}

//addPredefinedItem adds pre-defined items to the database
func addPredefinedItem() {
	item1 := &Item{
		Id:              getId(),
		Name:            "Apple",
		Category:        "Fruit",
		Price:           10,
		Status:          "1",
		DeletionComment: "",
		Description:     "Juicy and Healthy",
	}
	item2 := &Item{
		Id:              getId(),
		Name:            "Orange",
		Category:        "Fruit",
		Price:           8,
		Status:          "1",
		DeletionComment: "",
		Description:     "Juicy and Healthy",
	}
	item3 := &Item{
		Id:              getId(),
		Name:            "Nike Hoodie",
		Category:        "Clothes",
		Price:           30.5,
		Status:          "1",
		DeletionComment: "",
		Description:     "Comfortable to wear, nice style",
	}
	item4 := &Item{
		Id:              getId(),
		Name:            "Coke",
		Category:        "Drinks",
		Price:           3,
		Status:          "1",
		DeletionComment: "",
		Description:     "Taste Good",
	}
	item5 := &Item{
		Id:              getId(),
		Name:            "Pillow",
		Category:        "Bedding",
		Price:           18.5,
		Status:          "1",
		DeletionComment: "",
		Description:     "Soft and Comfortable",
	}
	shopItemList = append(shopItemList, item1, item2, item3, item4, item5)
}

// checkEqualItem returns true if the given item satisfy all the given equality conditions, otherwise false
func checkEqualItem(equalCods map[string]string, item *Item) bool {
	equal := true
	if idStr, ok := equalCods["id"]; ok {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return false
		}
		equal = equal && item.Id == id
	}
	if name, ok := equalCods["name"]; ok {
		equal = equal && item.Name == name
	}
	if category, ok := equalCods["category"]; ok {
		equal = equal && item.Category == category
	}
	if status, ok := equalCods["status"]; ok {
		equal = equal && item.Status == status
	}
	return equal
}

// checkGreaterItem returns true if the given item satisfy all the given inequality(greater than) conditions
func checkGreaterItem(greaterCods map[string]string, item *Item) bool {
	greater := true
	if priceStr, ok := greaterCods["price"]; ok {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return false
		}
		greater = greater && item.Price >= price
	}
	return greater
}

// checkLessItem returns true if the given item satisfy all the given inequality(less than) conditions
func checkLessItem(lessCods map[string]string, item *Item) bool {
	less := true
	if priceStr, ok := lessCods["price"]; ok {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return false
		}
		less = less && item.Price <= price
	}
	return less
}

// SelectWhereItem select items from the database that satisfy all the given conditions
func SelectWhereItem(equalCods map[string]string, greaterCods map[string]string, lessCods map[string]string) ([]*Item, int64, error) {
	resItemList := make([]*Item, 0)
	return selectWhere(equalCods, greaterCods, lessCods, resItemList)
}

// InsertItem insert a new item to the database, assigned a globally unique id to the item
func InsertItem(item *Item) (int64, error) {
	item.Id = getId()
	shopItemList = append(shopItemList, item)
	return 1, nil
}

// UpdateItemStatusWhereId updates the status (deleted, undeleted) the item base on the given id
func UpdateItemStatusWhereId(id int64, targetStatus string, comment string) (int64, error) {
	// id strictly increasing, apply binary search
	l, r := 0, len(shopItemList)-1
	for {
		if l >= r {
			break
		}
		mid := (l + r + 1) >> 1
		if shopItemList[mid].Id <= id {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if shopItemList[r].Id != id {
		return 0, errors.New("error: item not found")
	}
	if shopItemList[r].Status == targetStatus {
		return 0, errors.New("error: invalid item status")
	}
	shopItemList[r].Status = targetStatus
	if targetStatus == "0" {
		// indicates this is a deletion, add deletion comments
		shopItemList[r].DeletionComment = comment
	} else {
		// indicates this is an undeletion, renew deletion comments to ""
		shopItemList[r].DeletionComment = ""
	}
	return 1, nil
}

// UpdateItemWhereId updates some  features of the item of the given id base on the passed editMap
func UpdateItemWhereId(id int64, editMap map[string]interface{}) (int64, error) {
	// id strictly increasing, apply binary search
	l, r := 0, len(shopItemList)-1
	for {
		if l >= r {
			break
		}
		mid := (l + r + 1) >> 1
		if shopItemList[mid].Id <= id {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if shopItemList[r].Id != id {
		return 0, errors.New("error: item not found")
	}
	item := shopItemList[r]
	affectedRows := int64(0)
	if name, ok := editMap["name"]; ok {
		affectedRows++
		item.Name = name.(string)
	}
	if price, ok := editMap["price"]; ok {
		affectedRows++
		item.Price = price.(float64)
	}
	if description, ok := editMap["description"]; ok {
		affectedRows++
		item.Description = description.(string)
	}
	if category, ok := editMap["category"]; ok {
		affectedRows++
		item.Category = category.(string)
	}

	return affectedRows, nil
}
