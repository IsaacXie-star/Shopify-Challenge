package welcome_handler

import (
	"Shopify-Challenge/biz/constants"
	"Shopify-Challenge/biz/handler/base"
	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	base.ErrorHandler(c, nil, constants.RootPathErrorCode, constants.RootPathErrorMsg)
}
