package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeerCSV(t *testing.T) {
	// Test case 1: Valid CSV file
	lines, err := LeerCSV("../data/tests/valid.csv")
	assert.NoError(t, err, "no se esperaba un error")
	assert.NotNil(t, lines, "las líneas no deberían ser nulas")

	// Test case 2: Invalid CSV file
	lines, err = LeerCSV("../data/tests/invalid.csv")
	assert.Error(t, err, "se esperaba un error")
	assert.Nil(t, lines, "las líneas deberían ser nulas")
	assert.EqualError(t, err, "el CSV no contiene todos los campos necesarios", "mensaje de error inesperado")

	// Test case 3: Empty CSV file
	lines, err = LeerCSV("../data/tests/empty.csv")
	assert.Error(t, err, "se esperaba un error")
	assert.Nil(t, lines, "las líneas deberían ser nulas")
	assert.EqualError(t, err, "el CSV esta vacio", "mensaje de error inesperado")

	// Test case 4: Archivo no encontrado
	nombreArchivoInexistente := "archivo_no_existente.csv"
	lines, err = LeerCSV(nombreArchivoInexistente)
	assert.Error(t, err, "se esperaba un error")
	assert.Nil(t, lines, "las líneas deberían ser nulas")
	assert.EqualError(t, err, "open "+nombreArchivoInexistente+": no such file or directory", "mensaje de error inesperado")

	// Test case 5: Missing columns
	lines, err = LeerCSV("../data/tests/missing_columns.csv")
	assert.Error(t, err, "se esperaba un error")
	assert.Nil(t, lines, "las líneas deberían ser nulas")
	assert.EqualError(t, err, "el CSV no contiene todos los valores necesarios", "mensaje de error inesperado")

}

func TestExtraerDatosCSV(t *testing.T) {
	lines := [][]string{
		{"Nombre", "Temporada", "Partidos", "Puntos", "Equipo", "Asistencias", "Rebotes", "Tapones", "Robos", "Perdidas"},
		{"Kobe Bryant", "2009", "82", "2500", "LAL", "400", "500", "50", "100", "200"},
	}

	// Test case 1: Valid input
	expected := EstadisticasJugador{
		nombreApellidos: "Kobe Bryant",
		partidosJugados: 82,
		puntos:          2500,
		equipo:          "LAL",
		temporada:       2009,
		asistencias:     400,
		rebotes:         500,
		tapones:         50,
		robos:           100,
		perdidas:        200,
	}
	actual, err := extraerDatosCSV(lines, 1)
	assert.NoError(t, err, "error inesperado")
	assert.Equal(t, expected, actual, "resuultado inesperado")

	// Test case 2: Invalid input (non-integer values)
	lines[1][2] = "82a"
	lines[1][3] = "2500a"
	actual, err = extraerDatosCSV(lines, 1)
	assert.EqualError(t, err, "error al convertir los datos", "error inesperado")
	assert.Equal(t, EstadisticasJugador{}, actual, "resultado inesperado")

}

func TestObtenerIndiceColumna(t *testing.T) {
	headers := []string{"Nombre", "Temporada", "Partidos", "Puntos", "Equipo", "Asistencias", "Rebotes", "Tapones", "Robos", "Perdidas"}

	// Test case 1: Column exists
	expected := 2
	actual, err := obtenerIndiceColumna(headers, "Partidos")
	assert.NoError(t, err, "error inesperado")
	assert.Equal(t, expected, actual, "resultado inesperado")

	// Test case 2: Column does not exist
	_, err = obtenerIndiceColumna(headers, "Faltante")
	assert.Contains(t, err.Error(), "no se ha encontrado la columna", "error inesperado")
}

func TestExisteJugadorEpoca(t *testing.T) {
	epoca := Epoca{
		estadisticasJugadores: map[Clave]EstadisticasJugador{
			Clave{nombreApellidos: "Kobe Bryant", temporada: 2009}: EstadisticasJugador{},
		},
	}

	// Test case 1: Existing player in epoch
	clave1 := Clave{nombreApellidos: "Kobe Bryant", temporada: 2009}
	existe1 := existeJugadorEpoca(epoca, clave1)
	assert.True(t, existe1, "expected player to exist in epoch")

	// Test case 2: Non-existing player in epoch
	clave2 := Clave{nombreApellidos: "Michael Jordan", temporada: 1998}
	existe2 := existeJugadorEpoca(epoca, clave2)
	assert.False(t, existe2, "expected player to not exist in epoch")
}

func TestExisteEpoca(t *testing.T) {
	epocas := map[int]Epoca{
		2000: Epoca{},
	}

	// Test case 1: Existing epoch
	existe1 := existeEpoca(epocas, 2000)
	assert.True(t, existe1, "expected epoch to exist")

	// Test case 2: Non-existing epoch
	existe2 := existeEpoca(epocas, 2010)
	assert.False(t, existe2, "expected epoch to not exist")
}

func TestObtenerAñoInicioEpoca(t *testing.T) {
	// Test case 1: Season 2009
	expected1 := 2000
	actual1 := obtenerAñoInicioEpoca(2009)
	assert.Equal(t, expected1, actual1, "unexpected result")

	// Test case 2: Season 2016
	expected2 := 2010
	actual2 := obtenerAñoInicioEpoca(2016)
	assert.Equal(t, expected2, actual2, "unexpected result")

	// Test case 3: Season 1998
	expected3 := 1990
	actual3 := obtenerAñoInicioEpoca(1998)
	assert.Equal(t, expected3, actual3, "unexpected result")

	// Test case 4: Season 2000
	expected4 := 2000
	actual4 := obtenerAñoInicioEpoca(2000)
	assert.Equal(t, expected4, actual4, "unexpected result")
}

func TestCrearNuevaEpoca(t *testing.T) {
	// Test case 1: Starting year 2000
	expected1 := Epoca{
		fechaInicio:           2000,
		fechaFin:              2009,
		estadisticasJugadores: make(map[Clave]EstadisticasJugador),
	}
	actual1 := crearNuevaEpoca(2000)
	assert.Equal(t, expected1, actual1, "unexpected result")

	// Test case 2: Starting year 2010
	expected2 := Epoca{
		fechaInicio:           2010,
		fechaFin:              2019,
		estadisticasJugadores: make(map[Clave]EstadisticasJugador),
	}
	actual2 := crearNuevaEpoca(2010)
	assert.Equal(t, expected2, actual2, "unexpected result")

	// Test case 3: Starting year 1990
	expected3 := Epoca{
		fechaInicio:           1990,
		fechaFin:              1999,
		estadisticasJugadores: make(map[Clave]EstadisticasJugador),
	}
	actual3 := crearNuevaEpoca(1990)
	assert.Equal(t, expected3, actual3, "unexpected result")
}

func TestAñadeJugadorEpoca(t *testing.T) {
	epocas := map[int]Epoca{}
	clave := Clave{nombreApellidos: "Kobe Bryant", temporada: 2009}
	jugador := EstadisticasJugador{temporada: 2009, puntos: 100}

	// Test case 1: Existing epoch
	epocas[2000] = Epoca{
		estadisticasJugadores: map[Clave]EstadisticasJugador{
			clave: EstadisticasJugador{temporada: 2009, puntos: 100},
		},
	}

	añadeJugadorEpoca(epocas, clave, jugador)
	assert.Equal(t, 1, len(epocas), "unexpected number of epochs")
	assert.Equal(t, 1, len(epocas[2000].estadisticasJugadores), "unexpected number of players in epoch")
	assert.Equal(t, jugador, epocas[2000].estadisticasJugadores[clave], "unexpected player statistics")

	// Test case 2: Non-existing epoch
	clave2 := Clave{nombreApellidos: "Lebron James", temporada: 2016}
	jugador2 := EstadisticasJugador{temporada: 2016, puntos: 300}
	añadeJugadorEpoca(epocas, clave2, jugador2)
	assert.Equal(t, 2, len(epocas), "unexpected number of epochs")
	assert.Equal(t, 1, len(epocas[2010].estadisticasJugadores), "unexpected number of players in epoch")
	assert.Equal(t, jugador2, epocas[2010].estadisticasJugadores[clave2], "unexpected player statistics")

	// Test case 3: Player already exists in epoch
	añadeJugadorEpoca(epocas, clave2, jugador2)
	assert.Equal(t, 2, len(epocas), "unexpected number of epochs")
	assert.Equal(t, 1, len(epocas[2010].estadisticasJugadores), "unexpected number of players in epoch")
	assert.Equal(t, jugador2, epocas[2010].estadisticasJugadores[clave2], "unexpected player statistics")
}
func TestAñadeEstadisticas(t *testing.T) {
	lines := [][]string{
		{"Nombre", "Temporada", "Partidos", "Puntos", "Equipo", "Asistencias", "Rebotes", "Tapones", "Robos", "Perdidas"},
		{"Kobe Bryant", "2009", "82", "2500", "LAL", "400", "500", "50", "100", "200"},
		{"Lebron James", "2016", "82", "2000", "CLE", "600", "700", "70", "150", "250"},
	}

	// Test case 1: Valid input
	expected := map[int]Epoca{
		2000: Epoca{
			fechaInicio: 2000,
			fechaFin:    2009,
			estadisticasJugadores: map[Clave]EstadisticasJugador{
				Clave{nombreApellidos: "Kobe Bryant", temporada: 2009}: EstadisticasJugador{
					nombreApellidos: "Kobe Bryant",
					partidosJugados: 82,
					puntos:          2500,
					equipo:          "LAL",
					temporada:       2009,
					asistencias:     400,
					rebotes:         500,
					tapones:         50,
					robos:           100,
					perdidas:        200,
				},
			},
		},
		2010: Epoca{
			fechaInicio: 2010,
			fechaFin:    2019,
			estadisticasJugadores: map[Clave]EstadisticasJugador{
				Clave{nombreApellidos: "Lebron James", temporada: 2016}: EstadisticasJugador{
					nombreApellidos: "Lebron James",
					partidosJugados: 82,
					puntos:          2000,
					equipo:          "CLE",
					temporada:       2016,
					asistencias:     600,
					rebotes:         700,
					tapones:         70,
					robos:           150,
					perdidas:        250,
				},
			},
		},
	}

	actual := añadeEstadisticas(lines)
	assert.Equal(t, expected, actual, "unexpected result")
}
