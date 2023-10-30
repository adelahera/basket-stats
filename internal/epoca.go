package main

type Epoca struct {
	fechaInicio int
	fechaFinal  int
	jugadores   []Jugador
}

func NuevaEpoca(fechaInicio int, fechaFinal int, jugadores []Jugador) Epoca {
	epoca := Epoca{
		fechaInicio: fechaInicio,
		fechaFinal:  fechaFinal,
		jugadores:   jugadores,
	}
	return epoca
}
