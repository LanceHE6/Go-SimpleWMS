package model

import "time"

type VerificationCode struct {
	Uid       string    `gorm:"primary_key" json:"uid"`
	Email     string    `json:"bindEmail"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	Used      bool      `json:"used"`
}
