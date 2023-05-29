package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bnallapeta/logo-revelio/pkg/model"
	"github.com/bnallapeta/logo-revelio/pkg/store"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// handlers.go

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

func CheckAnswerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var submittedAnswer model.Answer
		if err := c.BindJSON(&submittedAnswer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println("submittedAnswer.LogoName: ", submittedAnswer.LogoName)
		correctLogoName, exists := store.CheckAnswer(submittedAnswer.LogoName, store.Logodata.Logos)
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logo name"})
			return
		}

		log.Println("correctLogoName: ", correctLogoName)
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

		log.Println("Score.FinalScore: ", Score.FinalScore)
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

		log.Println("coming till here?")
		c.JSON(http.StatusOK, gin.H{})
	}
}

func GetLogosHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logos := store.Logodata.Logos

		c.JSON(http.StatusOK, gin.H{
			"logos": logos,
		})
	}
}

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

// // GetUserScores fetches the user scores
// func GetUserScores(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Fetch all user scores from the database
// 		scores, err := store.GetScores(db)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Failed to fetch user scores",
// 			})
// 			return
// 		}
// 		// Respond with the list of user scores
// 		c.HTML(http.StatusOK, "userscores.html", gin.H{
// 			"scores": scores,
// 		})
// 	}
// }

// // GetGameSessions fetches the user scores
// func GetGameSessions(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Fetch all user scores from the database
// 		gamesessions, err := store.GetGameSessions(db)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Failed to fetch user scores",
// 			})
// 			return
// 		}
// 		// Respond with the list of user scores
// 		c.HTML(http.StatusOK, "gamesessions.html", gin.H{
// 			"gamesessions": gamesessions,
// 		})
// 	}
// }

// // SubmitResponse handles the submission of the user's response to the game.
// func SubmitResponse(db *gorm.DB) gin.HandlerFunc {
// 	type ResponseData struct {
// 		LogoName string `json:"logoname" binding:"required"`
// 		Answer   string `json:"answer" binding:"required"`
// 	}

// 	return func(c *gin.Context) {
// 		// Retrieve the session token from the cookie
// 		sessionToken, err := c.Cookie("session_token")
// 		log.Println("Session token: ", sessionToken)
// 		if err != nil {
// 			// Session token not found or invalid, handle the error accordingly
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid session",
// 			})
// 			log.Println("error 1: ", err)
// 			return
// 		}

// 		// Retrieve the session data from the session store
// 		sessionData := store.GetSessionData(sessionToken)
// 		if sessionData == nil {
// 			// Session data not found or expired, handle the error accordingly
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid session",
// 			})
// 			return
// 		}

// 		// Retrieve the user ID, game session ID, and logo ID from the session data
// 		userID := sessionData.UserID
// 		log.Println("userID that I retrieved from the session: ", userID)

// 		var data ResponseData
// 		if err := c.ShouldBindJSON(&data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Invalid request payload",
// 			})
// 			// TODO: handle this better. This is occuring due to empty textbox and user clicks submit
// 			log.Println("is it here?", err)
// 			return
// 		}

// 		// Check if the user's response is correct or not
// 		correctAnswer, ok := store.GetCorrectAnswer(data.LogoName, store.Logodata.Logos)
// 		if !ok {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Invalid logo name",
// 			})
// 			log.Println("or here?")
// 			return
// 		}

// 		isCorrect := strings.EqualFold(strings.ToLower(data.Answer), strings.ToLower(correctAnswer))

// 		// Get the user by name
// 		user, err := store.GetUserByID(db, userID)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Failed to get user",
// 			})
// 			return
// 		}

// 		// Get the current game session for the user
// 		gameSession, err := store.GetCurrentGameSession(db, user.ID)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Failed to get current game session",
// 			})
// 			return
// 		}

// 		// Create a new UserResponse entry
// 		userResponse := model.UserResponse{
// 			GameSessionID: gameSession.ID,
// 			LogoID:        0, // Set the appropriate LogoID
// 			UserResponse:  data.Answer,
// 			Correct:       isCorrect,
// 		}
// 		err = store.CreateUserResponse(db, &userResponse)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Failed to create user response",
// 			})
// 			return
// 		}

// 		// Update the user's score if the response is correct
// 		if isCorrect {
// 			err = store.UpdateUserScore(db, user.ID, 1)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"error": "Failed to update user score",
// 				})
// 				return
// 			}
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Response submitted",
// 		})
// 	}
// }
