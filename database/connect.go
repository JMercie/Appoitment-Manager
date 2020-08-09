package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JMercie/appointment-manager/config"
	"github.com/JMercie/appointment-manager/tables"
	"github.com/jinzhu/gorm"
)

// InitDatabase start the db
func InitDatabase() {

	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DBConn, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASS"), config.Config("DB_NAME")))
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")

	DBConn.AutoMigrate(&tables.User{}, &tables.Cliente{}, &tables.Empleado{}, &tables.Servicio{}, &tables.Turnos{})
	log.Println("Database Migrated")
}

// "host=localhost port=5432 user=postgres dbname=imperiogold sslmode=disable password="+pass
