package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
}

type RegisterReponse struct {
}

func Register(c *gin.Context) {
	c.String(http.StatusOK, "注册接口 ")

}
