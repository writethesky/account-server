package model

import (
	"account-server/repository/entity"
	"encoding/json"
)

type Account struct {
	Model
	UserID uint64             `json:"user_id"`
	Title  string             `gorm:"type:varchar(30)" json:"title"`
	Type   entity.AccountType `gorm:"type:tinyint" json:"type"`
	Data   interface{}        `gorm:"type:json" json:"data"`
}

func ToAccount(account entity.Account) (accountModel Account, err error) {
	accountModel = Account{
		Model: Model{
			ID:        account.Model.ID,
			CreatedAt: account.Model.CreatedAt.UnixMilli(),
			UpdatedAt: account.Model.UpdatedAt.UnixMilli(),
			DeletedAt: account.Model.DeletedAt.Time.UnixMilli(),
		},
		UserID: account.UserID,
		Title:  account.Title,
		Type:   account.Type,
	}

	switch account.Type {
	case entity.AccountTypeNormal:
		accountNormal := new(entity.AccountNormal)
		err = json.Unmarshal([]byte(account.Data), accountNormal)
		if nil != err {
			return
		}
		accountModel.Data = accountNormal
	case entity.AccountTypeEmail:
		accountEmail := new(entity.AccountEmail)
		err = json.Unmarshal([]byte(account.Data), accountEmail)
		if nil != err {
			return
		}
		accountModel.Data = accountEmail
	case entity.AccountTypeSSH:
		accountSSH := new(entity.AccountSSH)
		err = json.Unmarshal([]byte(account.Data), accountSSH)
		if nil != err {
			return
		}
		accountModel.Data = accountSSH
	}

	return
}

func ToAccountList(accountList []entity.Account) (list []Account, err error) {
	list = make([]Account, 0, len(accountList))
	for _, account := range accountList {
		accountModel, err := ToAccount(account)
		if nil != err {
			return nil, err
		}
		list = append(list, accountModel)
	}
	return
}
