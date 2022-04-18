package middleware

import (
	"Shopify-Challenge/biz/constants"
	"Shopify-Challenge/biz/handler/base"
	"github.com/gin-gonic/gin"
	"log"
)

//PanicRecover is a middleware used to recover panic and avoid app from crashing
func PanicRecover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic! Please check the code carefully")
			base.ErrorHandler(c, nil, constants.InternalErrorCode, constants.InternalErrorMsg)
			c.Abort()
		}
	}()
	c.Next()
}
