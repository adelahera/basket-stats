package main

type Clave struct {
	nombreApellidos string
	temporada       int
}

type Epoca struct {
	fechaInicio           int
	fechaFin              int
	estadisticasJugadores map[Clave]EstadisticasJugador
}
