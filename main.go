package main

import (
	"fmt"
	"github.com/artylark/todo-go-api/infrastructure"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API called")
	})

	db := infrastructure.Connect()
	fmt.Println(db)

	e.Logger.Fatal(e.Start(":8080"))
}
