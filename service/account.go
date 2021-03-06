package service

import (
	"account-server/repository/dao"
	"account-server/repository/entity"
	"account-server/repository/model"
	"encoding/json"
	"errors"
)

type CreateAccountInput struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Type        entity.AccountType `json:"type"`
	Data        interface{}        `json:"data"`
}

func CreateAccount(userID int64, input CreateAccountInput) (accountModel model.Account, err error) {
	dataBytes, err := json.Marshal(input.Data)
	if nil != err {
		return
	}

	if "" == input.Title {
		return accountModel, errors.New("the title cannot be empty")
	}
	switch input.Type {
	case entity.AccountTypeNormal:
		err = checkAccountNormal(dataBytes)
	case entity.AccountTypeEmail:
		err = checkAccountEmail(dataBytes)
	case entity.AccountTypeSSH:
		err = checkAccountSSH(dataBytes)
	}
	if nil != err {
		return
	}

	account := &entity.Account{
		UserID:      uint64(userID),
		Title:       input.Title,
		Description: input.Description,
		Type:        input.Type,
		Data:        string(dataBytes),
	}
	err = dao.CreateAccount(account)
	if nil != err {
		return
	}
	return model.ToAccount(*account)
}

func GetAccountInfo(userID uint64, accountID uint) (account model.Account, err error) {
	accountEntity, err := dao.GetAccount(accountID)
	if nil != err {
		return
	}
	if accountEntity.UserID != userID {
		return account, errors.New("no permission")
	}
	return model.ToAccount(accountEntity)
}

func GetAccountList(userID int64) (list []model.Account, err error) {
	accountList, err := dao.GetAccountList(userID)
	if nil != err {
		return
	}
	list, err = model.ToAccountList(accountList)
	return
}

func DeleteAccount(userID, accountID int64) (err error) {
	return dao.DeleteAccount(userID, accountID)
}

func ModifyAccount(userID int64, accountID int64, input CreateAccountInput) (account model.Account, err error) {
	dataBytes, err := json.Marshal(input.Data)
	if nil != err {
		return
	}
	err = dao.UpdateAccount(userID, accountID, entity.Account{
		Title:       input.Title,
		Type:        input.Type,
		Description: input.Description,
		Data:        string(dataBytes),
	})
	if nil != err {
		return
	}
	accountEntity, err := dao.GetAccount(uint(accountID))
	if nil != err {
		return
	}
	return model.ToAccount(accountEntity)
}

func checkAccountNormal(dataBytes []byte) (err error) {
	accountNormal := new(model.AccountNormal)
	err = json.Unmarshal(dataBytes, accountNormal)
	if nil != err {
		return
	}

	if "" == accountNormal.Account || "" == accountNormal.Password {
		return errors.New("the account or password cannot be empty")
	}
	return
}

func checkAccountEmail(dataBytes []byte) (err error) {
	accountEmail := new(model.AccountEmail)
	err = json.Unmarshal(dataBytes, accountEmail)
	if nil != err {
		return
	}

	if "" == accountEmail.Account || "" == accountEmail.Password || "" == accountEmail.LoginAddress {
		return errors.New("the account or password or login address cannot be empty")
	}
	return
}

func checkAccountSSH(dataBytes []byte) (err error) {
	accountSSH := new(model.AccountSSH)
	err = json.Unmarshal(dataBytes, accountSSH)
	if nil != err {
		return
	}

	if "" == accountSSH.User || "" == accountSSH.Address {
		return errors.New("the user or address cannot be empty")
	}
	return
}
