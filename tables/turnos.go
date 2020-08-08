package tables

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Turnos turnos
type Turnos struct {
	gorm.Model

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
