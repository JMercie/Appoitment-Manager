package tables

// Cliente cliente
type Cliente struct {
	ID int

	Nombre string

	Telefono int

	Turnos []Turnos
}
