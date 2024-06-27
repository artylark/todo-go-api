package main

import (
	"github.com/artylark/todo-go-api/infrastructure/datastore"
	"github.com/artylark/todo-go-api/infrastructure/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := datastore.Connect()

	router.Router(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
