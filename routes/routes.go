package routes

import (
	"github.com/JMercie/appointment-manager/auth"
	"github.com/JMercie/appointment-manager/handler"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	//middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	//Auth
	authO := api.Group("/auth")
	authO.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", auth.Protected(), handler.UpdateUser)
	user.Delete("/:id", auth.Protected(), handler.DeleteUser)

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
