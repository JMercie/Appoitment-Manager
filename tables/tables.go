package tables

import (
	"encoding/json"
	"log"
	"strconv"
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

	Empleado string

	Cliente string

	Fecha *time.Time

	Hora *time.Time

	Precio int

	Asistio bool

	ServicioID uint64

	EmpleadoID uint64

	ClienteID uint64
}

func GetEmpleados(c *fiber.Ctx) {

	db := database.DBConn

	var empleado []empleado

	db.Table("empleado").Find(&empleado)

	c.JSON(&empleado)
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

//this method shoudl perform this query SELECT  empleado.nombre as empleado, turnos.fecha, turnos.hora, turnos.cliente_id as cliente FROM turnos INNER JOIN empleado ON empleado.id = turnos.empleado_id INNER JOIN cliente ON cliente.id = turnos.cliente_id;
func GetTurnosWithEmpleado(c *fiber.Ctx) {

	id := c.Params("id")

	db := database.DBConn

	var turnos []turnos

	db.Table("turnos").Select("_id, empleado.nombre as empleado, turnos.fecha, turnos.hora, cliente.nombre as cliente, servicio.precio as precio").
		Joins("JOIN empleado ON empleado.id = turnos.empleado_id").
		Joins("JOIN cliente ON cliente.id = turnos.cliente_id").
		Joins("JOIN servicio ON servicio.id = turnos.servicio_id").
		Where("empleado_id = ?", id).
		Scan(&turnos)

	c.JSON(turnos)
}

func GetTurnosWithCliente(c *fiber.Ctx) {

	id := c.Params("id")

	db := database.DBConn

	var turnos []turnos

	db.Table("turnos").Select("_id, empleado.nombre as empleado, turnos.fecha, turnos.hora, cliente.nombre as cliente, servicio.precio as precio").
		Joins("JOIN empleado ON empleado.id = turnos.empleado_id").
		Joins("JOIN cliente ON cliente.id = turnos.cliente_id").
		Joins("JOIN servicio ON servicio.id = turnos.servicio_id").
		Where("cliente_id = ?", id).
		Scan(&turnos)

	c.JSON(turnos)
}

func UpdateTurnos(c *fiber.Ctx) {

	id := c.Params("id")
	tf := c.Params("tf")

	db := database.DBConn

	var turnos []turnos

	if err := db.Model(&turnos).Where("_id = ?", id).Update("asistio", tf).Error; err != nil {
		log.Fatal("not possible to update")
	}
	log.Printf("succesfuly update turno: %s", id)
}

func DeleteTurnos(c *fiber.Ctx) {

	id := c.Params("id")

	db := database.DBConn

	var turnos []turnos

	if err := db.Delete(&turnos, id).Error; err != nil {
		log.Fatal("not possible to update")
	}
	log.Printf("succesfuly update turno: %s", id)
}

func CreateTurnos(c *fiber.Ctx) {

	fecha := c.Params("date")
	hora := c.Params("time")
	eid := c.Params("eid") // id del empleado asignado
	cid := c.Params("cid") // id del cliente que posee el turno
	sid := c.Params("sid") // id del servicio seleccionado

	var servicio servicio
	db := database.DBConn

	//	 este bloque parseo los strings de los id a uint para poder usarlos en la creacion del registro
	eID, err := strconv.ParseUint(eid, 10, 64)
	if err != nil {
		log.Print(err)
	}
	cID, err := strconv.ParseUint(cid, 10, 64)
	if err != nil {
		log.Print(err)
	}
	sID, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		log.Print(err)
	}

	// este bloque seteo un formato para la hora y fecha del registro
	layoutFecha := "2006-01-02"
	layoutHora := "15:04"

	fechaTotime, err := time.Parse(layoutFecha, fecha)
	if err != nil {
		log.Print(err)
	}
	horaTotime, err := time.Parse(layoutHora, hora)
	if err != nil {
		log.Print(err)
	}

	// busco los servicios cuyo id es el solicitado
	db.Table("servicio").Where("id =?", sid).Find(&servicio)

	jsonstring, err := json.Marshal(servicio)
	if err != nil {
		log.Println(err)
	}

	// convierto a json los servicios para asi traer el valor de precio de cada uno y asigarno al registro
	byteArray := json.Unmarshal([]byte(jsonstring), &servicio)
	if byteArray != nil {
		log.Println(byteArray)
	}

	turno := turnos{
		Fecha:      &fechaTotime,
		Hora:       &horaTotime,
		EmpleadoID: eID,
		ClienteID:  cID,
		ServicioID: sID,
		Precio:     servicio.Precio,
	}

	db.Create(&turno)
	log.Printf("record not created? %t", db.NewRecord(turno))
}

//// aun trabajando en esta funcionalidad, hasta ahora estoy parseando mal las fechas o no las paso bien como parametro y
