package dao

import (
	"account-server/internal"
	"account-server/repository/entity"
	"account-server/repository/model"
	"encoding/json"
)

func CreateAccount(account entity.Account) (accountModel model.Account, err error) {
	err = internal.DB.Create(&account).Error
	if nil != err {
		return
	}

	return GetAccount(account.ID)
}

func GetAccount(id uint) (account model.Account, err error) {
	accountEntity := new(entity.Account)

	err = internal.DB.Where("id=?", id).First(&accountEntity).Error
	if nil != err {
		return
	}

	account = model.Account{
		Model: model.Model{
			ID:        accountEntity.Model.ID,
			CreatedAt: accountEntity.Model.CreatedAt,
			UpdatedAt: accountEntity.Model.UpdatedAt,
			DeletedAt: accountEntity.Model.DeletedAt,
		},
		Title: accountEntity.Title,
		Type:  accountEntity.Type,
	}
	switch accountEntity.Type {
	case entity.AccountTypeNormal:
		accountNormal := new(entity.AccountNormal)
		err = json.Unmarshal([]byte(accountEntity.Data), accountNormal)
		if nil != err {
			return
		}
		account.Data = accountNormal
	case entity.AccountTypeEmail:
		accountEmail := new(entity.AccountEmail)
		err = json.Unmarshal([]byte(accountEntity.Data), accountEmail)
		if nil != err {
			return
		}
		account.Data = accountEmail
	case entity.AccountTypeSSH:
		accountSSH := new(entity.AccountSSH)
		err = json.Unmarshal([]byte(accountEntity.Data), accountSSH)
		if nil != err {
			return
		}
		account.Data = accountSSH
	}

	return
}
