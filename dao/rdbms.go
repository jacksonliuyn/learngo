package dao

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type UsersInfo struct {
	gorm.Model
	Username     string `json:"username" gorm:"column:username;unique"`
	OldPassword  string `json:"old_password" gorm:"column:old_password"`
	NewPassword  string `json:"new_password" gorm:"column:new_password"`
	UserToken    string `json:"user_token" gorm:"column:usertoken"`
	TokenExpired int    `json:"token_expired" gorm:"column:token_expired"`
}

// TableName 通过构造UsersInfo方法，返回带模式的表
func (UsersInfo) TableName() string {
	return "users_info"
}

func PgCliTest() {
	dsn := "host=localhost user=test password=123456 dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	// 禁用复数形式
	err1 := db.AutoMigrate(&UsersInfo{})
	if err != nil {
		fmt.Println(err1)
	}

	userInfo := UsersInfo{
		Username:    "张三",
		NewPassword: "123456",
	}
	db.Create(&userInfo)
}
