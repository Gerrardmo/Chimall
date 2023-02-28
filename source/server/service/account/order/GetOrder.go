package serviceOrder

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetOrderRequest struct {
}

type GetOrderReponse struct {
}

func GetOrder(c *gin.Context) {
	//绑定字符
	req := ListRequest{}
	c.ShouldBindJSON(&req)

	//输出字符
	res := ListReponse{}
	c.JSON(http.StatusOK, res)
}
