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
	UserID     uint `gorm:"index"`
	Score      int
	AchievedAt time.Time
	User       User `gorm:"foreignkey:UserID"`
}

type Answer struct {
	gorm.Model
	UserID     uint `gorm:"index"`
	User       User `gorm:"foreignkey:UserID"`
	LogoName   string
	UserAnswer string
}
