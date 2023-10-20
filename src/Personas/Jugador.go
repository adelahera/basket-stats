package Personas

type Jugador struct {
	nombre          string
	apellidos       string
	equipos         []string
	partidosJugados int
	puntos          int
	asistencias     int
	rebotes         int
	tapones         int
	robos           int
	perdidas        int
}

func NuevoJugador(nombre string, apellidos string, equipos []string,
partidosJugados int, puntos int, asistencias int, rebotes int, tapones int,
robos int, perdidas int) Jugador {
	jugador := Jugador{
		nombre:          nombre,
		apellidos:       apellidos,
		equipos:         equipos,
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
