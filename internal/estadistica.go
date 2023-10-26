package Estadisticas

import e "src/Agrupacion"

type Estadistica struct {
	Etapas []e.Etapa
}

func NuevaEstadistica(etapas []e.Etapa) Estadistica {
	estadistica := Estadistica{
		Etapas: etapas,
	}
	return estadistica
}
