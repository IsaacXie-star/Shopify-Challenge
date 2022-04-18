package base

import (
	"Shopify-Challenge/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessHandlerWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.NewResponseWithData(data))
}

func SuccessHandlerWithMsg(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, model.NewResponseWithMsg(data, msg))
}

func SuccessHandlerWithDataTotal(c *gin.Context, data interface{}, total int64) {
	c.JSON(http.StatusOK, model.NewResponseWithTotal(data, total))
}

func SuccessHandlerWithDataPageSize(c *gin.Context, data interface{}, page int64, size int64, total int64) {
	c.JSON(http.StatusOK, model.NewResponseWithPageSize(data, page, size, total))
}

func ErrorHandler(c *gin.Context, data interface{}, errCode int64, errMsg string) {
	c.JSON(http.StatusOK, model.NewResponseWithErrCodeAndMsg(data, errCode, errMsg))
}
