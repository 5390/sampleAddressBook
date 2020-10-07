package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleAddresBook/services"

	"github.com/labstack/echo/v4"
)

func SearchAddress(c echo.Context) error {
	searchKey := c.Param("SEARCHKEY")
	addresBook, _ := services.AddressBookService().SearchAddress(searchKey)
	buf, err := json.Marshal(addresBook)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	response := string(buf)
	return c.String(http.StatusOK, response)
}
