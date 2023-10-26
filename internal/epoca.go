package main

import (
	Jugador "src/Personas"
	"time"
)

type Etapa struct {
	FechaInicio time.Time
	FechaFinal  time.Time
	Jugadores   []Jugador.Jugador
}

func NuevaEtapa(fechaInicio time.Time, fechaFinal time.Time, jugadores []Jugador.Jugador) Etapa {
	etapa := Etapa{
		FechaInicio: fechaInicio,
		FechaFinal:  fechaFinal,
		Jugadores:   jugadores,
	}
	return etapa
}
