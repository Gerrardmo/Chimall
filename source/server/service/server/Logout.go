package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutRequest struct {
}

type LogoutReponse struct {
}

func Logout(c *gin.Context) {
	c.String(http.StatusOK, "登出接口 ")

}
