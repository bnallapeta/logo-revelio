package handler

import (
	"net/http"

	"github.com/bnallapeta/logo-revelio/pkg/store"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUserHandler handles the creation of a new user and starts the game
func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")

		newUser, err := store.AddUserAndStartSession(db, name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start a new game session"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"userId": newUser.ID})
	}
}

// GameHandler fetches the logos and passes them to game.html
func GameHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		logos, err := store.GetLogos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "game.html", gin.H{
			"userId": userId,
			"logos":  logos,
		})
	}
}
