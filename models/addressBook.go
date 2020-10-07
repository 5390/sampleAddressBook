package models

type AddressBook struct {
	Id           int    `db:"id" json:"id"`
	FirstName    string `db:"first_name" json:"firstName"`
	LastName     string `db:"last_name" json:"lastName"`
	AddressLine1 string `db:"address_line_1" json:"addresLine1"`
	AddressLine2 string `db:"address_line_2" json:"addresLine2"`
	State        string `db:"state" json:"state"`
	ZipCode      string `db:"zipcode" json:"zipCode"`
}
