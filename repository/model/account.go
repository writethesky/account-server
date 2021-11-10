package model

import (
	"account-server/repository/entity"
)

type Account struct {
	Model
	Title string             `gorm:"type:varchar(30)" json:"title"`
	Type  entity.AccountType `gorm:"type:tinyint" json:"type"`
	Data  interface{}        `gorm:"type:json" json:"data"`
}
