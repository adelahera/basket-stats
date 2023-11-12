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

func añadeEstadisticas(lines [][]string) map[int]Epoca {

	epocas := make(map[int]Epoca)

	for i := 1950; i < 2020; i += 10 {
		estadisticasJugadores := make(map[Clave]EstadisticasJugador)
		for j, line := range lines {
			if j == 0 || line[1] == "" {
				continue
			}

			temporada, err := strconv.Atoi(line[1])

			if err != nil {
				log.Fatal(err)
			}

			if temporada >= i && temporada < i+10 {
				nombre := line[2]
				clave := Clave{
					nombreApellidos: nombre,
					temporada:       temporada,
				}

				if _, existe := estadisticasJugadores[clave]; existe {
					continue
				}

				partidos, _ := strconv.Atoi(line[6])
				pts, _ := strconv.Atoi(line[52])
				eq := line[5]
				asis, _ := strconv.Atoi(line[47])
				reb, _ := strconv.Atoi(line[46])
				tap, _ := strconv.Atoi(line[49])
				rob, _ := strconv.Atoi(line[48])
				perd, _ := strconv.Atoi(line[50])

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

				estadisticasJugadores[clave] = jugador

			}
		}

		epocas[i] = Epoca{
			fechaInicio:           strconv.Itoa(i),
			fechaFin:              strconv.Itoa(i + 9),
			estadisticasJugadores: estadisticasJugadores,
		}
	}
	return epocas
}
