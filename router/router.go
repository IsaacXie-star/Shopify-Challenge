package router

import (
	"Shopify-Challenge/biz/handler/shop_item_handler"
	"Shopify-Challenge/biz/handler/welcome_handler"
	"Shopify-Challenge/middleware"
	"github.com/gin-gonic/gin"
)

// RouteRegister register self-defined routers
func RouteRegister(r *gin.Engine) {

	shop := r.Group("/shop/inventory")

	shop.Use(middleware.PanicRecover) // panic recover middleware, avoid app from crashing

	shop.GET("/", welcome_handler.Welcome)
	item := shop.Group("/item")
	{
		item.GET("/get_item", shop_item_handler.QueryShopItemList)
		item.GET("/get_item_details", shop_item_handler.QueryShopItemById)
		item.POST("/edit_item", shop_item_handler.EditShopItem)
		item.POST("/add_item", shop_item_handler.AddShopItem)
		item.POST("/delete_item", shop_item_handler.DeleteShopItem)
		item.POST("/undelete_item", shop_item_handler.UnDeleteShopItem)
	}
}
