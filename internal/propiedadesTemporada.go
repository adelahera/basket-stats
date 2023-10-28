package main

type PropiedadesTemporada struct {
	puntos      int
	asistencias int
	rebotes     int
	tapones     int
	robos       int
	perdidas    int
}

func propiedadesTemporada(puntos int, asistencias int, rebotes int, tapones int, robos int, perdidas int) PropiedadesTemporada {
	propiedadesTemporada := PropiedadesTemporada{
		puntos:      puntos,
		asistencias: asistencias,
		rebotes:     rebotes,
		tapones:     tapones,
		robos:       robos,
		perdidas:    perdidas,
	}
	return propiedadesTemporada
}

