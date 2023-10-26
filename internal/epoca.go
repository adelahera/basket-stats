package main



type Etapa struct {
	fechaInicio int
	fechaFinal  int
	jugadores   []Jugador
}

func NuevaEtapa(fechaInicio int, fechaFinal int, jugadores []Jugador) Etapa {
	etapa := Etapa{
		fechaInicio: fechaInicio,
		fechaFinal:  fechaFinal,
		jugadores:   jugadores,
	}
	return etapa
}
