package main

type Estadistica struct {
	epocas []Epoca
}

func NuevaEstadistica(epocas []Epoca) Estadistica {
	estadistica := Estadistica{
		epocas: epocas,
	}
	return estadistica
}

func normalizaEpoca(epoca Epoca) Epoca {
	return NuevaEpoca(epoca.fechaInicio, epoca.fechaFinal, epoca.jugadores, epoca.reglas, true)
}
