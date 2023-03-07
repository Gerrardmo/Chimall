package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/component/config"
	"server/logic/orm/dal"
	accountService "server/service/account/account"
	serviceOrder "server/service/account/order"
	"server/service/server"
)

func main() {
	g := gin.New()
	//

	const dsn = "chimall:password@tcp(localhost:3306)/chimall?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn))
	//换成配置化  从config文件读取
	db, err := gorm.Open(mysql.Open(config.Config.GetString("mysql.dsn")))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	dal.SetDefault(db)

	acc, err := dal.Account.Where(dal.Account.ID.Eq("a")).First()

	fmt.Println(acc)
	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout", server.Logout)
	//

	account := g.Group("/account")
	order := account.Group("/order")

	order.POST("List", serviceOrder.List)

	account.POST("/profile", accountService.Profile)

	g.Run(":9999")
}
