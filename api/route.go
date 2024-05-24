package api

import (
	"cute/internal/game"
	login_record "cute/internal/login"
	"cute/internal/order"
	"cute/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPingRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", func(c *gin.Context) {
		data, err := user.FindAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.POST("/users", func(c *gin.Context) {
		var newUser user.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := user.CreateUser(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newUser)
	})
}

func RegisterGameRoutes(r *gin.Engine) {
	r.GET("/games", func(c *gin.Context) {
		data, err := game.FindAllGames()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})
}

func RegisterGameTypeRoutes(r *gin.Engine) {
	r.GET("/game-type", func(c *gin.Context) {
		data, err := game.FindAllGameTypes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})
}

func RegisterLoginRecordRoutes(r *gin.Engine) {
	r.GET("/login-record", func(c *gin.Context) {
		data, err := login_record.FindAllLoginRecords()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.POST("/login-record", func(c *gin.Context) {
		var login login_record.LoginRecord
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := login_record.CreateLoginRecord(&login); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, login)
	})
}

func RegisterOrderRoutes(r *gin.Engine) {
	r.GET("/orders", func(c *gin.Context) {
		orders, err := order.FindAllPurchaseOrders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	})
	r.POST("/orders", func(c *gin.Context) {
		var purchaseOrder order.PurchaseOrder
		if err := c.ShouldBindJSON(&purchaseOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := order.CreatePurchaseOrder(&purchaseOrder); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, purchaseOrder)
	})

}
