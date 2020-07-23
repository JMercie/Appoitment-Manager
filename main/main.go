package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Empleado struct {
	gorm.Model

	_Id int

	Nombre string

	Turnos []Turnos
}

type Servicio struct {
	gorm.Model

	_Id int

	Nombre string

	Precio int

	Turnos []Turnos
}

type Cliente struct {
	gorm.Model

	_Id int

	Nombre string

	Telefono int

	Fecha string

	Hora string

	Turnos []Turnos
}

type Turnos struct {
	gorm.Model

	_Id int

	Fecha string

	hora string

	Precio int

	Asistio bool

	Servicio_id []Servicio

	Empleado_id []Empleado

	Cliente_id []Cliente
}

var db *gorm.DB

var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=db port=5432 user=jmercie dbname=ImperioGold sslmode=disable password=solvay1017")

	if err != nil {
		panic("failed to connect do db")
	}

	defer db.Close()

	router.HandleFunc("/empleado/{id}", GetEmpleados).Methods("GET")

	router.HandleFunc("/cliente/{id}", GetClientes).Methods("GET")

	router.HandleFunc("/servicio/{id}", GetServicios).Methods("GET")

	router.HandleFunc("/turnos/{id}", GetTurnos).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func GetEmpleados(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var empleado Empleado

	db.First(&empleado, params["id"])

	json.NewEncoder(w).Encode(&empleado)

}

func GetClientes(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var cliente Cliente

	db.First(&cliente, params["id"])

	json.NewEncoder(w).Encode(&cliente)

}

func GetServicios(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var servicio Servicio

	db.First(&servicio, params["id"])

	json.NewEncoder(w).Encode(&servicio)
}

func GetTurnos(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var turnos Turnos

	var empleado []Empleado
	var cliente []Cliente
	var servicio []Servicio

	db.First(&empleado, params["id"])
	db.First(&cliente, params["id"])
	db.First(&servicio, params["id"])

	db.Model(&empleado).Related(&empleado)
	db.Model(&cliente).Related(&cliente)
	db.Model(&servicio).Related(&servicio)

	turnos.Empleado_id = empleado
	turnos.Servicio_id = servicio
	turnos.Cliente_id = cliente

	json.NewEncoder(w).Encode(&turnos)

}
