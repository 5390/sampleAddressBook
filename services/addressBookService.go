package services

import (
	"simpleAddresBook/dao"
	"simpleAddresBook/models"
)

type addressBookService struct{}

type AddressBookServiceIF interface {
	SearchAddress(key string) ([]models.AddressBook, error)
}

func AddressBookService() AddressBookServiceIF {
	return &addressBookService{}
}

func (self *addressBookService) SearchAddress(key string) ([]models.AddressBook, error) {
	addresBook, addressBookError := dao.AddressBookServiceDao().SearchAddress(key)
	return addresBook, addressBookError
}
