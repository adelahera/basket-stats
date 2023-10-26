package main

type Jugador struct {
	nombreApellidos	string
	equipo          string
	temporada       int
	partidosJugados int
	puntos          int
	asistencias     int
	rebotes         int
	tapones         int
	robos           int
	perdidas        int
}

func NuevoJugador(nombreApellidos string, equipo string, temporada int,
	partidosJugados int, puntos int, asistencias int, rebotes int, tapones int,
	robos int, perdidas int) Jugador {
	jugador := Jugador{
		nombreApellidos: nombreApellidos,
		equipo:          equipo,
		temporada:       temporada,
		partidosJugados: partidosJugados,
		puntos:          puntos,
		asistencias:     asistencias,
		rebotes:         rebotes,
		tapones:         tapones,
		robos:           robos,
		perdidas:        perdidas,
	}
	return jugador
}
