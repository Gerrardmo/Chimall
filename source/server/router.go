package main

import (
	"github.com/gin-gonic/gin"
	accountService "server/service/account/account"
	serviceOrder "server/service/account/order"
	"server/service/h"
	"server/service/server"
)

func Router(g *gin.Engine) {
	v1 := g.Group("api/v1")

	{
		v1.POST("/login", server.Login)
		v1.POST("/register", server.Register)
		v1.POST("/logout", server.Logout)
	}

	account := g.Group("/account")
	account.Use(h.Auth())
	order := account.Group("/order")

	order.POST("List", serviceOrder.List)

	account.POST("/profile", accountService.Profile)
}
