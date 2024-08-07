package model

type User struct {
	Base
	Uid        string `gorm:"primary_key;index" json:"uid"`
	Account    string `json:"account"`
	Password   string `json:"-"`
	Permission int    `json:"permission"`
	Nickname   string `json:"nickname"`
	Phone      string `gorm:"type:varchar(100);default:''" json:"phone"`
	Email      string `gorm:"type:varchar(100);default:''" json:"email"`
	Token      string `gorm:"type:varchar(255);default:''" json:"-"`
	SessionID  string `gorm:"type:varchar(255);default:''" json:"-"`
}
