package tables

import "github.com/jinzhu/gorm"

// Empleado empleado
type Empleado struct {
	gorm.Model

	Nombre string

	Turnos []Turnos `gorm:"foreignkey:EmpleadoID;association_foreignkey:ID"`
}
