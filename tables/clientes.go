package tables

import "github.com/jinzhu/gorm"

// Cliente cliente
type Cliente struct {
	gorm.Model

	Nombre string

	Telefono int64

	Turnos []Turnos `gorm:"foreignkey:ClienteID;association_foreignkey:ID"`
}
