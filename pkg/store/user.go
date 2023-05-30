package store

import (
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/bnallapeta/logo-revelio/pkg/model"
	"gorm.io/gorm"
)

// AddUserAndStartSession creates a new user in the db
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

// CheckAnswer validates if the user provided answer is correct or not
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

// UpdateScore updates the score of the user based on their responses
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

// Helper functions to get data from db
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
