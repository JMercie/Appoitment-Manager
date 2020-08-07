package handler

import (
	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/tables"

	"github.com/gofiber/fiber"
)

func GetClientes(c *fiber.Ctx) {

	db := database.DBConn

	var clientes []tables.Cliente

	db.Table("cliente").Find(&clientes)

	c.JSON(&clientes)
}
