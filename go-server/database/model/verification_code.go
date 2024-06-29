package model

import "time"

type VerificationCode struct {
	Email     string    `gorm:"primary_key" json:"email"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	Used      bool      `json:"used"`
}
