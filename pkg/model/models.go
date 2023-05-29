package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100)"`
	Scores  []Score
	Answers []Answer
}

type Score struct {
	gorm.Model
	UserID     uint `gorm:"index"` // foreign key
	Score      int
	AchievedAt time.Time
	User       User
}

type Answer struct {
	gorm.Model
	UserID     uint `gorm:"index"` // foreign key
	User       User
	LogoName   string
	UserAnswer string
}
