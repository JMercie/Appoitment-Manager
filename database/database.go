package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // need this to work
)

var (
	// DBConn is use to initialize the db in the handlers
	DBConn *gorm.DB
)
