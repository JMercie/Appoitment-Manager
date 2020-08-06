package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/tables"
)

func initDatabase() {

	var err error

	pass := os.Getenv("IMPERIOGOLD")

	database.DBConn, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=imperiogold sslmode=disable password="+pass)
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

	app.Get("/ganancias", tables.TotalEarning)

	app.Post("/asistio/:id/:tf", tables.Asistio)

	app.Post("/createturno/:fecha/:hora/:eid/:cid/:sid", tables.CreateTurnos)

	app.Delete("/deleteturno/:id", tables.DeleteTurnos)
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
