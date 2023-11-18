package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type ComparacionEstadistica struct {
	epocas   map[int]Epoca
	factores []float64
}

var camposDeseados = []string{"Nombre", "Equipo", "Temporada", "Partidos", "Puntos", "Asistencias", "Rebotes", "Tapones", "Robos", "Perdidas"}

func compruebaCampoCSV(line []string, campo string) bool {

	for _, campoCSV := range line {
		if campoCSV == campo {
			return true
		}
	}

	return false
}

func compruebaTodosCamposCSV(line []string, campos []string) bool {

	for _, campo := range campos {
		if !compruebaCampoCSV(line, campo) {
			return false
		}
	}

	return true
}

func LeerCSV(nombre string) ([][]string, error) {

	file, err := os.Open(nombre)
	if err != nil {
		return nil, errors.New("open " + nombre + ": no such file or directory")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("el CSV no contiene todos los valores necesarios")
	}

	if len(lines) == 0 {
		return nil, errors.New("el CSV esta vacio")
	}

	if !compruebaTodosCamposCSV(lines[0], camposDeseados) {
		return nil, errors.New("el CSV no contiene todos los campos necesarios")
	}

	return lines, err
}

func obtenerIndiceColumna(line []string, nombreColumna string) (int, error) {
	for i, columna := range line {
		if columna == nombreColumna {
			return i, nil
		}
	}

	return -1, errors.New("no se ha encontrado la columna " + nombreColumna)
}

func extraerDatosCSV(lines [][]string, fila int) (EstadisticasJugador, error) {
	indicePlayer, _ := obtenerIndiceColumna(lines[0], "Nombre")
	indiceYear, _ := obtenerIndiceColumna(lines[0], "Temporada")
	indiceG, _ := obtenerIndiceColumna(lines[0], "Partidos")
	indicePTS, _ := obtenerIndiceColumna(lines[0], "Puntos")
	indiceTm, _ := obtenerIndiceColumna(lines[0], "Equipo")
	indiceAST, _ := obtenerIndiceColumna(lines[0], "Asistencias")
	indiceTRB, _ := obtenerIndiceColumna(lines[0], "Rebotes")
	indiceBLK, _ := obtenerIndiceColumna(lines[0], "Tapones")
	indiceSTL, _ := obtenerIndiceColumna(lines[0], "Robos")
	indiceTOV, _ := obtenerIndiceColumna(lines[0], "Perdidas")

	nombre := lines[fila][indicePlayer]
	temporada, err := strconv.Atoi(lines[fila][indiceYear])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	partidos, err := strconv.Atoi(lines[fila][indiceG])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	pts, err := strconv.Atoi(lines[fila][indicePTS])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	eq := lines[fila][indiceTm]
	asis, err := strconv.Atoi(lines[fila][indiceAST])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	reb, err := strconv.Atoi(lines[fila][indiceTRB])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	tap, err := strconv.Atoi(lines[fila][indiceBLK])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	rob, err := strconv.Atoi(lines[fila][indiceSTL])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	perd, err := strconv.Atoi(lines[fila][indiceTOV])
	if err != nil {
		return EstadisticasJugador{}, errors.New("error al convertir los datos")
	}

	jugador := EstadisticasJugador{
		nombreApellidos: nombre,
		partidosJugados: partidos,
		puntos:          pts,
		equipo:          eq,
		temporada:       temporada,
		asistencias:     asis,
		rebotes:         reb,
		tapones:         tap,
		robos:           rob,
		perdidas:        perd,
	}

	return jugador, nil
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

func añadeJugadorEpoca(epocas map[int]Epoca, clave Clave, jugador EstadisticasJugador) {
	inicioEpoca := obtenerAñoInicioEpoca(jugador.temporada)
	if existeEpoca(epocas, inicioEpoca) {
		if existeJugadorEpoca(epocas[inicioEpoca], clave) {
			return
		} else {
			epocas[inicioEpoca].estadisticasJugadores[clave] = jugador
		}
	} else {
		epoca := crearNuevaEpoca(inicioEpoca)
		epoca.estadisticasJugadores[clave] = jugador
		epocas[inicioEpoca] = epoca
	}
}

func añadeEstadisticas(lines [][]string) map[int]Epoca {

	epocas := make(map[int]Epoca)

	for j, line := range lines {
		if j == 0 || line[1] == "" {
			continue
		}

		jugador, _ := extraerDatosCSV(lines, j)
		clave := Clave{
			nombreApellidos: jugador.nombreApellidos,
			temporada:       jugador.temporada,
		}

		añadeJugadorEpoca(epocas, clave, jugador)

	}

	return epocas
}
