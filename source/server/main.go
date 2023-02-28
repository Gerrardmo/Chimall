package main

import (
	"github.com/gin-gonic/gin"
	accountservice "server/service/account/account"
	serviceOrder "server/service/account/order"
	"server/service/server"
)

func main() {
	g := gin.New()

	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout", server.Logout)
	//

	account := g.Group("/account")
	order := account.Group("/order")

	order.POST("List", serviceOrder.List)

	account.POST("/profile", accountservice.Profile)

	g.Run(":9999")
}
