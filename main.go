package main

import (
	"Shopify-Challenge/db"
	"Shopify-Challenge/router"
	"github.com/gin-gonic/gin"
)

func init() {
	db.Init()
}

func main() {
	// init router
	r := gin.Default()
	// register router
	router.RouteRegister(r)

	r.Run("127.0.0.1:8080")
}
