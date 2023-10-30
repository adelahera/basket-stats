package main

type Jugador struct {
	nombreApellidos      string
	propiedadesTemporada PropiedadesTemporada
}

func NuevoJugador(nombreApellidos string, equipo string, temporada int,
	partidosJugados int, propiedadesTemporada PropiedadesTemporada) Jugador {
	jugador := Jugador{
		nombreApellidos:      nombreApellidos,
		propiedadesTemporada: propiedadesTemporada,
	}
	return jugador
}
