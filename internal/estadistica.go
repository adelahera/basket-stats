package main

type Estadistica struct {
	epocas   []Epoca
	factores []float64
}

func NuevaEstadistica(epocas []Epoca) Estadistica {
	estadistica := Estadistica{
		epocas: epocas,
	}
	return estadistica
}

func normalizaEpoca(epoca Epoca) Epoca {
	return NuevaEpoca(epoca.fechaInicio, epoca.fechaFinal, epoca.jugadores)
}
