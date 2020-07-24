package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	Turnos []turnos
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

	hora *time.Time

	Precio int

	Asistio bool

	ServicioID []servicio

	EmpleadoID []empleado

	ClienteID []cliente
}

var db *gorm.DB

var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=imperiogold sslmode=disable password=postgres")

	if err != nil {
		panic("failed to connect do db")
	}

	defer db.Close()

	router.HandleFunc("/empleado", GetEmpleados).Methods("GET")

	router.HandleFunc("/cliente", GetClientes).Methods("GET")

	router.HandleFunc("/servicio", GetServicios).Methods("GET")

	router.HandleFunc("/turnos", GetTurnos).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func GetEmpleados(w http.ResponseWriter, r *http.Request) {

	var empleado []empleado

	db.Table("empleado").Find(&empleado)

	json.NewEncoder(w).Encode(&empleado)
}

func GetClientes(w http.ResponseWriter, r *http.Request) {

	var clientes []cliente

	db.Table("cliente").Find(&clientes)

	json.NewEncoder(w).Encode(&clientes)
}

func GetServicios(w http.ResponseWriter, r *http.Request) {

	var servicios []servicio

	db.Table("servicio").Find(&servicios)

	json.NewEncoder(w).Encode(&servicios)

}

func GetTurnos(w http.ResponseWriter, r *http.Request) {

	var turnos []turnos

	db.Find(&turnos)

	json.NewEncoder(w).Encode(&turnos)
}
