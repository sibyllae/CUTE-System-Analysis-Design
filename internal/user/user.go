// internal/app/user/user.go
package user

import (
	"cute/database"
	"log"
	"time"
)

// 定义 User 模型
type User struct {
	UserID           int       `gorm:"primaryKey;autoIncrement"`
	Username         string    `gorm:"unique;size:50;not null"`
	Email            string    `gorm:"unique;size:100;not null"`
	Password         string    `gorm:"size:100;not null"`
	RegistrationDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// 指定 User 模型的表名为 'user'
func (User) TableName() string {
	return "User"
}

func AutoMigrate() {
	database.DB.AutoMigrate(&User{})
}

func CreateUser(user *User) error {
	result := database.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	log.Println("User created successfully")
	return nil
}

func FindAllUsers() ([]User, error) {
	var users []User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Found %d users\n", len(users))
	return users, nil
}
