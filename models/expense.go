package models

import "time"

type Expense struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Date        time.Time `json:"date" gorm:"not null;default:CURRENT_TIMESTAMP"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
}
