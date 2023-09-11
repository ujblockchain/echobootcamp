package pages

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func FormContext(c echo.Context) error {
	//get user product id input
	productID := c.FormValue("product_id")
	//get user product quantity input
	quantity := c.FormValue("quantity")
	//get user product color input
	cornColor := c.FormValue("corn_color")

	//print user form inputs
	fmt.Println(productID)
	fmt.Println(quantity)
	fmt.Println(cornColor)

	return nil
}
