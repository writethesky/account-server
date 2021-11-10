package dao

import (
	"account-server/internal"
	"account-server/repository/entity"
)

func CreateAccount(account *entity.Account) (err error) {
	return internal.DB.Create(account).Error
}

func GetAccount(id uint) (account entity.Account, err error) {
	err = internal.DB.Where("id=?", id).First(&account).Error
	return
}

func GetAccountList(userID int64) (list []entity.Account, err error) {
	err = internal.DB.Where("user_id=?", userID).Order("updated_at desc").Find(&list).Error
	return
}
