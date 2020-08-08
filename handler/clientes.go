package handler

import (
	"log"
	"strconv"

	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/tables"

	"github.com/gofiber/fiber"
)

// GetClientes query all clients in the db
func GetClientes(c *fiber.Ctx) {

	db := database.DBConn

	var clientes []tables.Cliente

	db.Table("cliente").Find(&clientes)

	c.JSON(&clientes)
}

// CreateClientes post a new client
func CreateClientes(c *fiber.Ctx) {
	name := c.Params("name")
	phone := c.Params("phone")

	phoneToInt, err := strconv.ParseInt(phone, 10, 64)
	if err != nil {
		log.Print(err)
	}

	db := database.DBConn

	var turnos []tables.Turnos

	if err := db.Find(&turnos).Scan(&turnos).Error; err != nil {
		log.Printf("couldn't get turnos %s", err)
	}

	client := tables.Cliente{
		Nombre:   name,
		Telefono: phoneToInt,
	}

	db.Create(&client)

	log.Printf("record not created? %t", db.NewRecord(client))
	c.JSON(fiber.Map{"status": "success", "message": "client created", "data": client})

}
