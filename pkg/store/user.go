package store

import (
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/bnallapeta/logo-revelio/pkg/model"
	"gorm.io/gorm"
)

func AddUserAndStartSession(db *gorm.DB, name string) (*model.User, error) {
	user := &model.User{
		Name: name,
	}
	// Add the user to the database
	result := db.Create(user)

	// Error handling
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func CheckAnswer(logoName string, logosMap map[string]string) (string, bool) {
	correctAnswerPath, ok := logosMap[logoName]
	log.Println("correctAnswerPath: ", correctAnswerPath)

	if ok {
		// Extract the logo name from the full path
		correctLogoName := filepath.Base(correctAnswerPath)
		// Strip off the file extension
		correctLogoName = strings.TrimSuffix(correctLogoName, filepath.Ext(correctLogoName))
		return correctLogoName, true
	}
	return "", false
}

func UpdateUserScore(db *gorm.DB, userID uint) error {
	// Get the user
	user := &model.User{}
	db.First(user, userID)

	// Create new score entry
	newScore := &model.Score{
		UserID:     userID,
		Score:      1, // Score starts at 1 for each correct answer
		AchievedAt: time.Now(),
	}

	// Save new score entry
	result := db.Create(newScore)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateScore(db *gorm.DB, userID uint, finalScore int) error {
	log.Println(finalScore)
	// Get the user
	user := &model.User{}
	if err := db.First(user, userID).Error; err != nil {
		return err
	}

	// Create new score entry
	newScore := &model.Score{
		UserID:     userID,
		Score:      finalScore,
		AchievedAt: time.Now(),
		User:       *user,
	}

	// Save new score entry
	result := db.Create(newScore)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllUsers(db *gorm.DB) ([]model.User, error) {
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetUserScores(db *gorm.DB) ([]model.Score, error) {
	var scores []model.Score
	result := db.Find(&scores)
	if result.Error != nil {
		return nil, result.Error
	}

	return scores, nil
}

func GetTopTenScores(db *gorm.DB) ([]model.Score, error) {
	var scores []model.Score
	result := db.Order("score desc").Limit(10).Find(&scores)
	if result.Error != nil {
		return nil, result.Error
	}

	return scores, nil
}

// // CreateUser creates a new user in the database.
// func CreateUser(db *gorm.DB, name string) (uint, error) {
// 	user := User{Name: name}
// 	result := db.Create(&user)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return user.ID, nil
// }

// // CreateUserResponse creates a new UserResponse entry in the database.
// func CreateUserResponse(db *gorm.DB, userResponse *model.UserResponse) error {
// 	result := db.Create(userResponse)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// // UpdateUserScore updates the user's score in the database.
// func UpdateUserScore(db *gorm.DB, userID uint, score int) error {
// 	// Retrieve the latest game session for the user
// 	var gameSession model.GameSession
// 	result := db.Where("user_id = ?", userID).Last(&gameSession)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	// Add the score to the total score of the latest game session
// 	gameSession.TotalScore += score

// 	// Save the updated game session
// 	result = db.Save(&gameSession)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }

// // GetUserByID retrieves a user by ID from the database.
// func GetUserByID(db *gorm.DB, userID uint) (*User, error) {
// 	user := &User{}

// 	if err := db.First(user, userID).Error; err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// // GetUserByName retrieves a user by name from the database.
// func GetUserByName(db *gorm.DB, name string) (*User, error) {
// 	user := &User{}

// 	result := db.Where("name = ?", name).First(user)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return nil, fmt.Errorf("user not found")
// 	}
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return user, nil
// }

// type UserScore struct {
// 	Username   string
// 	EndTime    time.Time
// 	TotalScore int
// }

// // GetScores retrieves all user scores from the database.
// func GetScores(db *gorm.DB) ([]UserScore, error) {
// 	var scores []UserScore
// 	result := db.Table("users").
// 		Select("users.name as username, game_sessions.end_time, game_sessions.total_score").
// 		Joins("JOIN game_sessions ON game_sessions.user_id = users.id").
// 		Scan(&scores)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return scores, nil
// }

// // GetGameSessions retrieves all game session details for users from the database.
// func GetGameSessions(db *gorm.DB) ([]model.GameSession, error) {
// 	var sessions []model.GameSession
// 	result := db.Preload("User").Find(&sessions)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return sessions, nil
// }

// // StartGame starts the game for a user
// func StartGame(db *gorm.DB, userID uint) error {
// 	var user model.User
// 	result := db.First(&user, userID)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	gameSession := model.GameSession{
// 		UserID:    user.ID,
// 		StartTime: time.Now(),
// 	}

// 	result = db.Create(&gameSession)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }
