package main

type Epoca struct {
	fechaInicio int
	fechaFinal  int
	jugadores   []Jugador
	reglas      ReglasEstilos
}

func NuevaEpoca(fechaInicio int, fechaFinal int, jugadores []Jugador, reglas ReglasEstilos, normaliza bool) Epoca {
	epoca := Epoca{
		fechaInicio: fechaInicio,
		fechaFinal:  fechaFinal,
		jugadores:   jugadores,
		reglas:      reglas,
	}
	return epoca
}
