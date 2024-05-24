// internal/app/order/order.go
package order

import (
	"cute/database"
	"cute/internal/game"
	"cute/internal/user"
	"log"
	"time"
)

// 定义 PurchaseOrder 模型
type PurchaseOrder struct {
	OrderID   int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"index"`
	GameID    int       `gorm:"index"`
	Amount    float64   `gorm:"type:decimal(10,2)"`
	OrderDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User      user.User `gorm:"foreignKey:UserID" json:"-"`
	Game      game.Game `gorm:"foreignKey:GameID" json:"-"`
}

// 指定 PurchaseOrder 模型的表名为 'PurchaseOrder'
func (PurchaseOrder) TableName() string {
	return "PurchaseOrder"
}

func AutoMigrate() {
	database.DB.AutoMigrate(&PurchaseOrder{})
}

func CreatePurchaseOrder(order *PurchaseOrder) error {
	result := database.DB.Create(order)
	if result.Error != nil {
		return result.Error
	}
	log.Println("PurchaseOrder created successfully")
	return nil
}

func FindAllPurchaseOrders() ([]PurchaseOrder, error) {
	var orders []PurchaseOrder
	result := database.DB.Preload("User").Preload("Game").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Found %d purchase orders\n", len(orders))
	return orders, nil
}
