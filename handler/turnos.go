package handler

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/JMercie/appointment-manager/database"
	"github.com/JMercie/appointment-manager/tables"
	"github.com/gofiber/fiber"
)

// GetTurnos query all the appointments in the db
func GetTurnos(c *fiber.Ctx) {

	db := database.DBConn

	var turnos []tables.Turnos

	db.Order("precio DESC").Find(&turnos)

	c.JSON(&turnos)
}

// GetTurnosWithEmpleado this method shoudl perform this query SELECT  empleado.nombre as empleado, turnos.fecha, turnos.hora, turnos.cliente_id as cliente FROM turnos INNER JOIN empleado ON empleado.id = turnos.empleado_id INNER JOIN cliente ON cliente.id = turnos.cliente_id;
func GetTurnosWithEmpleado(c *fiber.Ctx) {

	id := c.Params("id")

	db := database.DBConn

	var turnos []tables.Turnos

	db.Table("turnos").Select("turnos.id, empleado.nombre as empleado, turnos.fecha, turnos.hora, cliente.nombre as cliente, servicio.precio as precio").
		Joins("JOIN empleado ON empleado.id = turnos.empleado_id").
		Joins("JOIN cliente ON cliente.id = turnos.cliente_id").
		Joins("JOIN servicio ON servicio.id = turnos.servicio_id").
		Where("empleado_id = ?", id).
		Order("fecha DESC").
		Scan(&turnos)

	c.JSON(turnos)
}

// GetTurnosWithCliente brings all appointments for specific client id
func GetTurnosWithCliente(c *fiber.Ctx) {

	id := c.Params("id")

	db := database.DBConn

	var turnos []tables.Turnos

	db.Table("turnos").Select("turnos.id, empleado.nombre as empleado, turnos.fecha, turnos.hora, cliente.nombre as cliente, servicio.precio as precio").
		Joins("JOIN empleado ON empleado.id = turnos.empleado_id").
		Joins("JOIN cliente ON cliente.id = turnos.cliente_id").
		Joins("JOIN servicio ON servicio.id = turnos.servicio_id").
		Where("cliente_id = ?", id).
		Order("fecha DESC").
		Scan(&turnos)

	c.JSON(turnos)
}

// Asistio You update the turn to know if the person was able to go to the shop
func Asistio(c *fiber.Ctx) {

	id := c.Params("id")
	tf := c.Params("tf")

	db := database.DBConn

	var turnos []tables.Turnos

	if err := db.Model(&turnos).Where("id = ?", id).Update("asistio", tf).Error; err != nil {
		log.Fatal("not possible to update")
	}
	log.Printf("succesfuly update turno: %s", id)
}

// DeleteTurnos delete an appointment
func DeleteTurnos(c *fiber.Ctx) {

	id := c.Params("id")

	db := database.DBConn

	var turnos []tables.Turnos

	if err := db.Delete(&turnos, id).Error; err != nil {
		log.Fatal("not possible to update")
	}
	log.Printf("succesfuly delete turno: %s", id)
	c.JSON(fiber.Map{"status": "success", "message": "succesfuly delete turno", "data": turnos})
}

// CreateTurnos post an appointment
func CreateTurnos(c *fiber.Ctx) {

	fecha := c.Params("fecha")
	hora := c.Params("hora")
	eid := c.Params("eid") // id del empleado asignado
	cid := c.Params("cid") // id del cliente que posee el turno
	sid := c.Params("sid") // id del servicio seleccionado

	var servicio tables.Servicio
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
	log.Printf("esta es mi fecha %s", fecha)
	log.Printf("esta es mi hora %s", hora)
	// este bloque seteo un formato para la hora y fecha del registro
	layoutFecha := "2006-01-02"
	layoutHora := "15:04:05"

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

	turno := tables.Turnos{
		Fecha:      &fechaTotime,
		Hora:       &horaTotime,
		EmpleadoID: eID,
		ClienteID:  cID,
		ServicioID: sID,
		Precio:     servicio.Precio,
	}

	db.Create(&turno)

	log.Printf("record not created? %t", db.NewRecord(turno))
	c.JSON(fiber.Map{"status": "success", "message": "turno created", "data": turno})
}

// TotalEarning perform the following query: SELECT SUM(precio) FROM turnos;
func TotalEarning(c *fiber.Ctx) {

	db := database.DBConn
	var total []int
	var turnos []tables.Turnos

	if err := db.Find(&turnos).Where("asistio = ?", "true").Pluck("precio", &total).Error; err != nil {
		log.Printf("there was this err %s", err)
	}

	result := 0
	for _, v := range total {
		result += v
	}

	c.JSON(result)
}
