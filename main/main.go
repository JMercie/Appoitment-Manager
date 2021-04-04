package main

import (
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/handlebars"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/routes"
)

func main() {

	engine := handlebars.New("/Users/joseph/go/src/github.com/JMercie/Appoitment-Manager/public", ".hbs")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})
	app.Use(cors.New())

	database.InitDatabase()
	app.Use(middleware.Recover())

	routes.SetupRoutes(app)
	app.Listen(3000)

	log.Fatal(app.Listen(3000))

	defer database.DBConn.Close()
}
