package tables

type Cliente struct {
	ID int

	Nombre string

	Telefono int

	Turnos []Turnos
}
