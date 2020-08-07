package main

import (
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/routes"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("/root/go/src/github.com/JMercie/appointment-manager/public", ".html")

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
