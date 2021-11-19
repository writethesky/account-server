package entity

import "gorm.io/gorm"

type AccountType uint

const (
	AccountTypeNormal AccountType = iota
	AccountTypeEmail
	AccountTypeSSH
)

type Account struct {
	gorm.Model
	UserID      uint64      `json:"user_id"`
	Title       string      `gorm:"type:varchar(30)"`
	Description string      `gorm:"type:varchar(255)"`
	Type        AccountType `gorm:"type:tinyint"`
	Data        string      `gorm:"type:json"`
}
