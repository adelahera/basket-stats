package main


type Estadistica struct {
	etapas []Etapa
}

func NuevaEstadistica(etapas []Etapa) Estadistica {
	estadistica := Estadistica{
		etapas: etapas,
	}
	return estadistica
}
