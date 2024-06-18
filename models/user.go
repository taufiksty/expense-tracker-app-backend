package models

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}
