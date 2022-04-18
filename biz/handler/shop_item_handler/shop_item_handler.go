package shop_item_handler

import (
	"Shopify-Challenge/biz/constants"
	"Shopify-Challenge/biz/handler/base"
	"Shopify-Challenge/dao"
	"Shopify-Challenge/db"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//QueryShopItemList return all the items satisfy the conditions
func QueryShopItemList(c *gin.Context) {
	conditions := make(map[string]string)
	if name, ok := c.GetQuery("name"); ok {
		conditions["name"] = name
	}
	if category, ok := c.GetQuery("category"); ok {
		conditions["category"] = category
	}
	if minPrice, ok := c.GetQuery("min_price"); ok {
		conditions["min_price"] = minPrice
	}
	if maxPrice, ok := c.GetQuery("max_price"); ok {
		conditions["max_price"] = maxPrice
	}
	if status, ok := c.GetQuery("status"); ok {
		if status != "0" && status != "1" {
			base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
			return
		}
		conditions["status"] = status
	}
	itemListDao := dao.NewItemListDao()
	itemList, total, err := itemListDao.QueryShopItemList(conditions)
	if err != nil {
		log.Printf("[itemListDao.QueryShopItemList] error, err=%v", err)
		base.ErrorHandler(c, nil, constants.DbQueryErrorCode, constants.DbQueryErrorMsg)
		return
	}
	base.SuccessHandlerWithDataTotal(c, itemList, total)
}

// QueryShopItemById returns detail info of an item of a given id
func QueryShopItemById(c *gin.Context) {
	conditions := make(map[string]string)
	if id, ok := c.GetQuery("id"); ok {
		conditions["id"] = id
	} else {
		// id must be passed, this should be controlled by frontend
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	itemListDao := dao.NewItemListDao()
	itemList, total, err := itemListDao.QueryShopItemList(conditions)
	if err != nil {
		log.Printf("[itemListDao.QueryShopItemList] error, err=%v", err)
		base.ErrorHandler(c, nil, constants.DbQueryErrorCode, constants.DbQueryErrorMsg)
		return
	}
	if total != 1 {
		log.Println("QueryShopItemById error, total != 1")
		base.ErrorHandler(c, nil, constants.DbQueryErrorCode, constants.DbQueryErrorMsg)
		return
	}
	base.SuccessHandlerWithDataTotal(c, itemList[0], total)
}

// AddShopItem adds a new item to the inventory
func AddShopItem(c *gin.Context) {
	name, _ := c.GetPostForm("name")
	category, _ := c.GetPostForm("category")
	priceStr, _ := c.GetPostForm("price")
	description, _ := c.GetPostForm("description")

	if name == "" || priceStr == "" || category == "" {
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}

	item := &db.Item{
		Name:            name,
		Category:        category,
		Price:           price,
		Status:          "1",
		DeletionComment: "",
		Description:     description,
	}
	rowsAffected, err := db.InsertItem(item)
	if rowsAffected != 1 || err != nil {
		base.ErrorHandler(c, nil, constants.DbInsertErrorCode, constants.DbInsertErrorMsg)
		return
	}
	base.SuccessHandlerWithMsg(c, nil, constants.AddItemSuccessMsg)
}

// DeleteShopItem delete an item in the inventory, set status to "0", item should have status "1" before
func DeleteShopItem(c *gin.Context) {
	idStr, _ := c.GetPostForm("id")
	targetStatus, _ := c.GetPostForm("status")
	deletionComment, _ := c.GetPostForm("deletion_comment")

	if idStr == "" || targetStatus != "0" {
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("[DeleteShopItem] strconv.ParseInt(id) error, err=%v", err)
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	rowsAffected, err := db.UpdateItemStatusWhereId(id, targetStatus, deletionComment)
	if rowsAffected != 1 || err != nil {
		base.ErrorHandler(c, nil, constants.DbInsertErrorCode, constants.DbInsertErrorMsg)
		return
	}
	base.SuccessHandlerWithMsg(c, nil, constants.ModifyStatusSuccessMsg)
}

// UnDeleteShopItem undelete an item in the inventory, set status to "1", item should have status "0" before
func UnDeleteShopItem(c *gin.Context) {
	idStr, _ := c.GetPostForm("id")
	targetStatus, _ := c.GetPostForm("status")

	if idStr == "" || targetStatus != "1" {
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("[DeleteShopItem] strconv.ParseInt(id) error, err=%v", err)
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	rowsAffected, err := db.UpdateItemStatusWhereId(id, targetStatus, "")
	if rowsAffected != 1 || err != nil {
		base.ErrorHandler(c, nil, constants.DbInsertErrorCode, constants.DbInsertErrorMsg)
		return
	}
	base.SuccessHandlerWithMsg(c, nil, constants.ModifyStatusSuccessMsg)
}

// EditShopItem edit an item with the given id in the inventory base on what is passed
// fields can be edited includes name, description, price, category
func EditShopItem(c *gin.Context) {
	idStr, _ := c.GetPostForm("id")
	if idStr == "" {
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("[DeleteShopItem] strconv.ParseInt(id) error, err=%v", err)
		base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
		return
	}
	editMap := make(map[string]interface{})
	editCount := int64(0)

	if name, ok := c.GetPostForm("name"); ok {
		editMap["name"] = name
		editCount++
	}
	if category, ok := c.GetPostForm("category"); ok {
		editMap["category"] = category
		editCount++
	}
	if description, ok := c.GetPostForm("description"); ok {
		editMap["description"] = description
		editCount++
	}
	if priceStr, ok := c.GetPostForm("price"); ok {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			base.ErrorHandler(c, nil, constants.InvalidParamCode, constants.InvalidParamMsg)
			return
		}
		editMap["price"] = price
		editCount++
	}

	rowsAffected, err := db.UpdateItemWhereId(id, editMap)
	if rowsAffected != editCount || err != nil {
		base.ErrorHandler(c, nil, constants.DbInsertErrorCode, constants.DbInsertErrorMsg)
		return
	}
	base.SuccessHandlerWithMsg(c, nil, constants.EditSuccessMsg)
}
