package entity

import "gorm.io/gorm"

type AccountType uint

const (
	AccountTypeNormal AccountType = iota
	AccountTypeEmail
	AccountTypeSSH
)

type AccountNormal struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AccountEmail struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	LoginAddress string `json:"login_address"`
}

type AccountSSH struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
}

type Account struct {
	gorm.Model
	UserID uint64      `json:"user_id"`
	Title  string      `gorm:"type:varchar(30)" json:"title"`
	Type   AccountType `gorm:"type:tinyint" json:"type"`
	Data   string      `gorm:"type:json" json:"data"`
}
