package dao

import (
	"fmt"
	"simpleAddresBook/db"
	"simpleAddresBook/models"
)

type addressBookServiceDao struct{}

type AddressBookServiceDaoIF interface {
	SearchAddress(key string) ([]models.AddressBook, error)
}

func AddressBookServiceDao() AddressBookServiceDaoIF {
	return &addressBookServiceDao{}
}

func (self *addressBookServiceDao) SearchAddress(key string) ([]models.AddressBook, error) {
	addresBook := []models.AddressBook{}
	db, ConnectionErrs := db.SqlxConnect()
	if ConnectionErrs != nil {
		return addresBook, ConnectionErrs
	}
	sqlStatement := `SELECT * FROM address_book where first_name=$1 or last_name =$1`
	errMainSql := db.Select(&addresBook, sqlStatement, key)
	if errMainSql != nil {
		fmt.Println(errMainSql)
		return addresBook, errMainSql
	}
	defer db.Close()
	return addresBook, nil
}
