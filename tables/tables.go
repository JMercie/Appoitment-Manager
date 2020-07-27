package tables

import (
	"log"
	"time"

	"github.com/JMercie/appointment-manager/database"
	"github.com/gofiber/fiber"
)

type empleado struct {
	ID int

	Nombre string

	Turnos []turnos
}

type servicio struct {
	ID int

	Nombre string

	Precio int

	Turno turnos
}

type cliente struct {
	ID int

	Nombre string

	Telefono int

	Turnos []turnos
}

type turnos struct {
	ID int

	Fecha *time.Time

	Hora *time.Time

	Precio int

	Asistio bool

	ServicioID uint

	EmpleadoID uint

	ClienteID uint
}

func GetEmpleados(c *fiber.Ctx) {

	db := database.DBConn

	var empleado []empleado

	db.Table("empleado").Find(&empleado)

	c.JSON(&empleado)
}

//this method shoudl perform this query SELECT  empleado.nombre as empleado, turnos.fecha, turnos.hora, turnos.cliente_id as cliente FROM turnos INNER JOIN empleado ON empleado.id = turnos.empleado_id INNER JOIN cliente ON cliente.id = turnos.cliente_id;
func GetTurnoConEmpleado(c *fiber.Ctx) {

	//idEmpleado := c.Params("id")
	//idCliente := c.Params("idc")

	var err error

	db := database.DBConn

	var turno []turnos

	if err = db.Select("empleado.nombre as empleado, turnos.fecha, turnos.hora, turnos.cliente_id as cliente").Joins("JOIN empleado ON empleado.id = turnos.empleado_id").Joins("JOIN cliente ON cliente.id = turnos.cliente_id").Find(&turno).Error; err != nil {
		log.Fatal(err)
	}

	for _, ar := range turno {
		log.Print(ar)
		c.JSON(&turno)
	}

}

func GetClientes(c *fiber.Ctx) {

	db := database.DBConn

	var clientes []cliente

	db.Table("cliente").Find(&clientes)

	c.JSON(&clientes)
}

func GetServicios(c *fiber.Ctx) {

	db := database.DBConn

	var servicios []servicio

	db.Table("servicio").Find(&servicios)

	c.JSON(&servicios)

}

func GetTurnos(c *fiber.Ctx) {

	db := database.DBConn

	var turnos []turnos

	db.Find(&turnos)

	c.JSON(&turnos)
}
