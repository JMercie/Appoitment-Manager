package tables

import "github.com/jinzhu/gorm"

// Servicio servicio
type Servicio struct {
	gorm.Model

	Nombre string

	Precio int

	Turno Turnos `gorm:"foreignkey:ServicioID"`
}
