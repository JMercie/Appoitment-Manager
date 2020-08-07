package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/routes"
)

func main() {

	app := fiber.New()

	database.InitDatabase()
	app.Use(middleware.Recover())

	routes.SetupRoutes(app)
	app.Listen(3000)

	log.Fatal(app.Listen(3000))

	defer database.DBConn.Close()
}
