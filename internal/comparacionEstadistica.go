package main

import (
	"encoding/csv"
	"log"
	"os"
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
