package store

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	sessionTokenLength = 32
)

var sessionStore = make(map[string]SessionData)

// SessionData represents the session data associated with a session token
type SessionData struct {
	UserID uint
}

// generateSessionToken generates a session token associated with the given user ID.
func GenerateSessionToken(userID uint) (string, error) {
	tokenBytes := make([]byte, sessionTokenLength)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate session token: %w", err)
	}

	token := base64.URLEncoding.EncodeToString(tokenBytes)
	return fmt.Sprintf("%d-%s", userID, token), nil
}

// StoreSessionData stores the session data in the session store and returns the session token
// func StoreSessionData(data SessionData) string {
// 	sessionToken := uuid.New().String()
// 	sessionStore[sessionToken] = data
// 	return sessionToken
// }

// GetSessionData retrieves the session data from the session store based on the session token
func GetSessionData(sessionToken string) *SessionData {
	// Parse the session token to extract the userID
	var userID uint
	_, err := fmt.Sscanf(sessionToken, "%d", &userID)
	if err != nil {
		return nil
	}

	// Create a new session data with the retrieved userID
	data := &SessionData{
		UserID: userID,
	}

	return data
}

// // GetCurrentGameSession retrieves the current active game session for the given user ID.
// func GetCurrentGameSession(db *gorm.DB, userID uint) (*model.GameSession, error) {
// 	gameSession := &model.GameSession{}

// 	if err := db.Where("user_id = ? AND end_time IS NULL", userID).First(gameSession).Error; err != nil {
// 		return nil, err
// 	}

// 	return gameSession, nil
// }
