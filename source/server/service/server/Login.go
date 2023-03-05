package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
}

type LoginReponse struct {
}

/**

 */

func Login(c *gin.Context) {
	c.String(http.StatusOK, "登录接口 ?")

}
