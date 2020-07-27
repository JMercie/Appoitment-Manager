package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/tables"
)

func initDatabase() {

	var err error

	database.DBConn, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=imperiogold sslmode=disable password=postgres")
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
}

func setupRoutes(app *fiber.App) {
	app.Get("/empleado", tables.GetEmpleados)

	app.Get("/cliente", tables.GetClientes)

	app.Get("/servicio", tables.GetServicios)

	app.Get("/turnos", tables.GetTurnos)

	app.Get("/turnosconempleado/:id", tables.GetTurnosWithEmpleado)

	app.Get("/turnoscliente/:id", tables.GetTurnosWithCliente)

	app.Post("/asistio/:id/:tf", tables.UpdateTurnos)

	app.Post("/createturno/:fecha/:hora/:eid/:cid/:sid", tables.CreateTurnos)
}

func main() {
	app := fiber.New()

	initDatabase()
	app.Use(middleware.Recover())
	setupRoutes(app)
	app.Listen(3000)

	log.Fatal(app.Listen(3000))

	defer database.DBConn.Close()
}
