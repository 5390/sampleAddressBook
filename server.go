package main

import (
	"simpleAddresBook/db"
	"simpleAddresBook/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	_, err := db.DBConnect()
	if err != nil {
		panic(err)
	}
	e.Use(middleware.CORS())
	route.AddressBookRouteService(e)
	e.Start(":8900")
}
