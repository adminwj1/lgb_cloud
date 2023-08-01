package utils

import (
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db := global.APP.Configuration.Database
	dsn := db.UserName + ":" + db.PassWord + "@tcp(" + db.Address + ":" + db.Port + ")/" + db.DBname + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("dsn： ", dsn)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	open.AutoMigrate(&models.User{})
	return open
}
