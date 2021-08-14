package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	e := echo.New()

	e.GET("/people", models.AllPeople)
	e.GET("/people/:id", models.FindPersonByID)
	e.Logger.Fatal(e.Start(":5000"))
}
