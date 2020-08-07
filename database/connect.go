package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func InitDatabase() {

	var err error

	// p := config.Config("DB_PORT")
	// port, err := strconv.ParseUint(p, 10, 32)
	pass := os.Getenv("IMPERIOGOLD")

	DBConn, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=imperiogold sslmode=disable password="+pass) //fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("IMPERIOGOLD"), config.Config("DB_NAME")))
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
}

// "host=localhost port=5432 user=postgres dbname=imperiogold sslmode=disable password="+pass
