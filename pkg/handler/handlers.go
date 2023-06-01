package handler

import (
	"net/http"
	"strconv"

	"github.com/bnallapeta/logo-revelio/pkg/model"
	"github.com/bnallapeta/logo-revelio/pkg/store"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUserHandler handles the creation of a new user and starts the game
func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser, err := store.AddUserAndStartSession(db, user.Name)
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

// CheckAnswerHandler checks the user provided answer with the one available in the logo map
// and validates if its the correct answer
func CheckAnswerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var submittedAnswer model.Answer
		if err := c.BindJSON(&submittedAnswer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		correctLogoName, exists := store.CheckAnswer(submittedAnswer.LogoName, store.Logodata.Logos)
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logo name"})
			return
		}

		isCorrect := correctLogoName == submittedAnswer.UserAnswer

		c.JSON(http.StatusOK, gin.H{"correct": isCorrect})
	}
}

// UpdateFinalScoreHandler updates the final score for a user
func UpdateFinalScoreHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var Score struct {
			UserID     string `json:"userID"`
			FinalScore int    `json:"finalScore"`
		}

		if err := c.BindJSON(&Score); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, _ := strconv.Atoi(Score.UserID)

		// Update the score in the database
		err := store.UpdateScore(db, uint(userID), Score.FinalScore)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Score updated successfully"})
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
