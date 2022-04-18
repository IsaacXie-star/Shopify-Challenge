package shop_item_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// unit test file
// use gin test framework, test http request

type ShopItemSuite struct {
	suite.Suite
}

func (suite *ShopItemSuite) TestQueryShopItemList() {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/test", QueryShopItemList)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, req)
	resp := w.Result()
	respBody, _ := ioutil.ReadAll(resp.Body)
	suite.Equal(http.StatusOK, resp.StatusCode)
	fmt.Println(string(respBody))
}

func TestShopItemSuite(t *testing.T) {
	suite.Run(t, new(ShopItemSuite))
}
