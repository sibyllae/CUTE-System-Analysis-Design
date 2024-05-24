package login_record

import (
	"cute/database"
	"log"
	"time"
)

// 定义 LoginRecord 模型
type LoginRecord struct {
	RecordID  int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"index"`
	LoginTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// 指定 LoginRecord 模型的表名为 'LoginRecord'
func (LoginRecord) TableName() string {
	return "LoginRecord"
}

func AutoMigrate() {
	database.DB.AutoMigrate(&LoginRecord{})
}

func CreateLoginRecord(record *LoginRecord) error {
	result := database.DB.Create(record)
	if result.Error != nil {
		return result.Error
	}
	log.Println("LoginRecord created successfully")
	return nil
}

func FindAllLoginRecords() ([]LoginRecord, error) {
	var records []LoginRecord
	result := database.DB.Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Found %d login records\n", len(records))
	return records, nil
}
