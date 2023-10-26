package Personas

type Jugador struct {
	nombre          string
	apellidos       string
	equipo          string
	temporada       string
	partidosJugados int
	puntos          int
	asistencias     int
	rebotes         int
	tapones         int
	robos           int
	perdidas        int
}

func NuevoJugador(nombre string, apellidos string, equipo string, temporada string,
	partidosJugados int, puntos int, asistencias int, rebotes int, tapones int,
	robos int, perdidas int) Jugador {
	jugador := Jugador{
		nombre:          nombre,
		apellidos:       apellidos,
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
