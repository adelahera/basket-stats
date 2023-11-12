package main

type Clave struct {
	nombreApellidos string
	temporada       int
}

type Epoca struct {
	fechaInicio           string
	fechaFin              string
	estadisticasJugadores map[Clave]EstadisticasJugador
}
