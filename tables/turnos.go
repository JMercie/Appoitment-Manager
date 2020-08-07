package tables

import "time"

type Turnos struct {
	ID int

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
