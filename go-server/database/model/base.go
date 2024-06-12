package model

import "time"

type Time struct {
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type Base struct {
	ID uint `gorm:"primary_key;" json:"-"`
	Time
}
