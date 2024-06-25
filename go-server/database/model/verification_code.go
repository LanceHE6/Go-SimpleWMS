package model

import "time"

type VerificationCode struct {
	Id        string    `gorm:"primary_key" json:"id"`
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	Used      bool      `json:"used"`
}
