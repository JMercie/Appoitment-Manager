package routes

import (
	auth "github.com/JMercie/appointment-manager/auth"
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

	// Empleado
	empleado := app.Group("/empleado")
	empleado.Get("/", handler.GetEmpleados)

	// Cliente
	cliente := app.Group("/cliente")
	cliente.Get("/", handler.GetClientes)

	// Servicio
	servicio := app.Group("/servicio")
	servicio.Get("/", handler.GetServicios)

	// Turnos
	turnos := app.Group("/turnos")
	turnos.Get("/", handler.GetTurnos)
	turnos.Get("/turnosconempleado/:id", handler.GetTurnosWithEmpleado)
	turnos.Get("/turnoscliente/:id", handler.GetTurnosWithCliente)
	turnos.Get("/ganancias", handler.TotalEarning)
	turnos.Patch("/asistio/:id/:tf", handler.Asistio)
	turnos.Post("/createturno/:fecha/:hora/:eid/:cid/:sid", handler.CreateTurnos)
	turnos.Delete("/deleteturno/:id", handler.DeleteTurnos)
}
