// internal/app/game/game.go
package game

import (
	"cute/database"
	"log"
	"time"
)

// 定义 Game 模型
type Game struct {
	GameID      int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	ReleaseDate time.Time `gorm:"type:date"`
	GameTypeID  int       `gorm:"index"`
	GameType    GameType  `gorm:"foreignKey:GameTypeID"`
}

// 指定 User 模型的表名为 'user'
func (Game) TableName() string {
	return "Game"
}

// 定义 GameType 模型
type GameType struct {
	GameTypeID int    `gorm:"primaryKey;autoIncrement"`
	TypeName   string `gorm:"size:50;not null"`
}

// 指定 User 模型的表名为 'user'
func (GameType) TableName() string {
	return "GameType"
}

func AutoMigrate() {
	database.DB.AutoMigrate(&Game{}, &GameType{})
}

func CreateGame(game *Game) error {
	result := database.DB.Create(game)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Game created successfully")
	return nil
}

func FindAllGames() ([]Game, error) {
	var games []Game
	result := database.DB.Preload("GameType").Find(&games)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Found %d games\n", len(games))
	return games, nil
}

func CreateGameType(gameType *GameType) error {
	result := database.DB.Create(gameType)
	if result.Error != nil {
		return result.Error
	}
	log.Println("GameType created successfully")
	return nil
}

func FindAllGameTypes() ([]GameType, error) {
	var gameTypes []GameType
	result := database.DB.Find(&gameTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Found %d game types\n", len(gameTypes))
	return gameTypes, nil
}
