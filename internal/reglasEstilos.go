package main

type ReglaEstilo struct {
	aplica      bool
	anio        int
	nombreRegla string
	factor      []float64
}

type ReglasEstilos struct {
	totalReglasEstilos []ReglaEstilo
}

func NuevasReglas(reglasEstilos []ReglaEstilo) ReglasEstilos {
	totalReglasEstilos := ReglasEstilos{
		totalReglasEstilos: reglasEstilos,
	}
	return totalReglasEstilos
}
