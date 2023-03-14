package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/component/config"
	"server/logic/orm/dal"
)

func init() {
	//const dsn = "root:mcw123456.@tcp(47.115.134.176:3306)/chimall?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn))
	//换成配置化  从config文件读取
	db, err := gorm.Open(mysql.Open(config.Config.GetString("mysql.dsn")))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	dal.SetDefault(db)
}
