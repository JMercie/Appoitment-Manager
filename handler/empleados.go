package handler

import (
	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/tables"
	"github.com/gofiber/fiber"
)

// GetEmpleados query all employees in the db
func GetEmpleados(c *fiber.Ctx) {

	db := database.DBConn

	var empleado []tables.Empleado

	db.Table("empleado").Find(&empleado)

	c.JSON(&empleado)
}
