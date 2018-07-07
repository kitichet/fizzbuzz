package main

import (
	"net/http"
	
	"github.com/labstack/echo"
	"github.com/kitichet/fizzbuzz/fizzbuzz"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		
		return c.String(http.StatusOK, "Hello, World!")
	})


	e.GET("/fizzbuzz/:i", func(c echo.Context) error {
		i := c.Param("i")
		return c.String(http.StatusOK,fizzbuzz.Fizzbuzz(i))
	})
	e.Logger.Fatal(e.Start(":1323"))
}