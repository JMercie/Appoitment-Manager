package routes

import (
	"github.com/JMercie/appointment-manager/handler"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func SetupRoutes(app *fiber.App) {

	//middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	app.Get("/empleado", handler.GetEmpleados)
	app.Get("/cliente", handler.GetClientes)
	app.Get("/servicio", handler.GetServicios)
	app.Get("/turnos", handler.GetTurnos)
	app.Get("/turnosconempleado/:id", handler.GetTurnosWithEmpleado)
	app.Get("/turnoscliente/:id", handler.GetTurnosWithCliente)
	app.Get("/ganancias", handler.TotalEarning)

	app.Patch("/asistio/:id/:tf", handler.Asistio)

	app.Post("/createturno/:fecha/:hora/:eid/:cid/:sid", handler.CreateTurnos)

	app.Delete("/deleteturno/:id", handler.DeleteTurnos)
}
