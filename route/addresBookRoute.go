package route

import (
	"simpleAddresBook/controllers"

	"github.com/labstack/echo/v4"
)

func AddressBookRouteService(e *echo.Echo) {
	e.GET("/search/address/:SEARCHKEY", controllers.SearchAddress)
}
