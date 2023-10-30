package main

type PropiedadesTemporada struct {
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

func propiedadesTemporada(equipo string, temporada int, partidosJugados int, puntos int, asistencias int, rebotes int, tapones int, robos int, perdidas int) PropiedadesTemporada {
	propiedadesTemporada := PropiedadesTemporada{
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
	return propiedadesTemporada
}
