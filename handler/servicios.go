package handler

import (
	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/tables"
	"github.com/gofiber/fiber"
)

func GetServicios(c *fiber.Ctx) {

	db := database.DBConn

	var servicios []tables.Servicio

	db.Table("servicio").Find(&servicios)

	c.JSON(&servicios)

}
