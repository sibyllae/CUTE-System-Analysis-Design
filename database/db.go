// internal/db/db.go
package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// 建立 MySQL 连接字符串
	dsn := "root:123456@tcp(127.0.0.1:3306)/cute?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接到 MySQL
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}
