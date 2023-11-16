package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type ComparacionEstadistica struct {
	epocas   map[int]Epoca
	factores []float64
}

func LeerCSV(nombre string) ([][]string, error) {

	file, err := os.Open(nombre)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, err
}

func obtenerIndiceColumna(line []string, nombreColumna string) int {
	for i, columna := range line {
		if columna == nombreColumna {
			return i
		}
	}

	return -1
}

func extraerDatosCSV(lines [][]string, fila int) EstadisticasJugador {

	nombre := lines[fila][obtenerIndiceColumna(lines[0], "Player")]
	log.Println(obtenerIndiceColumna(lines[0], "Player"))
	temporada, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "Year")])
	partidos, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "G")])
	pts, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "PTS")])
	eq := lines[fila][obtenerIndiceColumna(lines[0], "Tm")]
	asis, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "AST")])
	reb, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "TRB")])
	tap, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "BLK")])
	rob, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "STL")])
	perd, _ := strconv.Atoi(lines[fila][obtenerIndiceColumna(lines[0], "TOV")])

	return jugador
}
func existeJugadorEpoca(epoca Epoca, clave Clave) bool {
	_, existe := epoca.estadisticasJugadores[clave]
	return existe
}

func existeEpoca(epocas map[int]Epoca, inicioEpoca int) bool {
	_, existe := epocas[inicioEpoca]
	return existe
}

func obtenerAñoInicioEpoca(temporada int) int {
	return (temporada / 10) * 10
}

func crearNuevaEpoca(inicioEpoca int) Epoca {
	epoca := Epoca{
		fechaInicio:           inicioEpoca,
		fechaFin:              inicioEpoca + 9,
		estadisticasJugadores: make(map[Clave]EstadisticasJugador),
	}

	return epoca
}
	return epocas
}
	return jugador
}
