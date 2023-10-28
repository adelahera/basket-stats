package main

type Jugador struct {
	nombreApellidos      string
	equipo               string
	temporada            int
	partidosJugados      int
	propiedadesTemporada PropiedadesTemporada
}

func NuevoJugador(nombreApellidos string, equipo string, temporada int,
	partidosJugados int, propiedadesTemporada PropiedadesTemporada) Jugador {
	jugador := Jugador{
		nombreApellidos:      nombreApellidos,
		equipo:               equipo,
		temporada:            temporada,
		partidosJugados:      partidosJugados,
		propiedadesTemporada: propiedadesTemporada,
	}
	return jugador
}


