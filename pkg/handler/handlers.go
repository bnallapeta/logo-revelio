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

// GetUsers handles the retrieval of all users.
func GetAllUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := store.GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "allusers.html", gin.H{
			"users": users,
		})
	}
}

// GetUserScores handles the retrieval of all users.
func GetUserScoresHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		scores, err := store.GetUserScores(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "userscores.html", gin.H{
			"scores": scores,
		})
	}
}

// GetTopTenScores handles the retrieval of all users.
func GetTopTenScoresHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		scores, err := store.GetTopTenScores(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "toptenscores.html", gin.H{
			"scores": scores,
		})
	}
}
