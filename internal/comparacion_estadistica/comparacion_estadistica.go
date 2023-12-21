package comparacion_estadistica

import (
	"encoding/csv"
	"errors"
	"math"
	"os"
	"strconv"

	"github.com/adelahera/basket-stats/internal/log"
)

var logger = log.GetLogger()

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
			logger.Debug().Msg("Campo " + campo + " no está en el CSV")
			return false
		}
	}

	return true
}

func LeerCSV(nombre string) ([][]string, error) {

	logger.Info().Msg("Leyendo CSV " + nombre)
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
	temporada, tempErr := strconv.Atoi(lines[fila][indiceYear])
	partidos, partidosErr := strconv.Atoi(lines[fila][indiceG])
	pts, ptsErr := strconv.Atoi(lines[fila][indicePTS])
	eq := lines[fila][indiceTm]
	asis, asisErr := strconv.Atoi(lines[fila][indiceAST])
	reb, rebErr := strconv.Atoi(lines[fila][indiceTRB])
	tap, tapErr := strconv.Atoi(lines[fila][indiceBLK])
	rob, robErR := strconv.Atoi(lines[fila][indiceSTL])
	perd, perdErr := strconv.Atoi(lines[fila][indiceTOV])

	if tempErr != nil || partidosErr != nil || ptsErr != nil || asisErr != nil || rebErr != nil || tapErr != nil || robErR != nil || perdErr != nil {
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
	if existe {
		logger.Debug().Msg("El jugador " + clave.nombreApellidos + " existe en la epoca " + strconv.Itoa(epoca.fechaInicio))
	} else {
		logger.Debug().Msg("El jugador " + clave.nombreApellidos + " no existe en la epoca " + strconv.Itoa(epoca.fechaInicio))
	}
	return existe
}

func existeEpoca(epocas map[int]Epoca, inicioEpoca int) bool {

	_, existe := epocas[inicioEpoca]
	if existe {
		logger.Debug().Msg("La epoca " + strconv.Itoa(inicioEpoca) + " existe")
	} else {
		logger.Debug().Msg("La epoca " + strconv.Itoa(inicioEpoca) + " no existe")
	}
	return existe
}

func obtenerAñoInicioEpoca(temporada int) int {

	logger.Info().Msg("Obteniendo año de inicio de la epoca " + strconv.Itoa(temporada))
	return (temporada / 10) * 10
}

func crearNuevaEpoca(inicioEpoca int) Epoca {

	logger.Info().Msg("Creando nueva epoca " + strconv.Itoa(inicioEpoca))
	epoca := Epoca{
		fechaInicio:           inicioEpoca,
		fechaFin:              inicioEpoca + 9,
		estadisticasJugadores: make(map[Clave]EstadisticasJugador),
	}

	return epoca
}

func añadeJugadorEpoca(epocas map[int]Epoca, clave Clave, jugador EstadisticasJugador) {

	logger.Info().Msg("Añadiendo jugador " + clave.nombreApellidos + " a la epoca " + strconv.Itoa(jugador.temporada))
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
		logger.Debug().Msg("Jugador " + jugador.nombreApellidos + " extraido")
		clave := Clave{
			nombreApellidos: jugador.nombreApellidos,
			temporada:       jugador.temporada,
		}

		añadeJugadorEpoca(epocas, clave, jugador)
	}

	return epocas
}

func calculaMediaEstadistica(epoca Epoca, estadistica string) float64 {

	logger.Info().Msg("Calculando media de " + estadistica + " en la epoca " + strconv.Itoa(epoca.fechaInicio))
	var suma float64
	for _, jugador := range epoca.estadisticasJugadores {
		switch estadistica {
		case "Partidos":
			suma += float64(jugador.partidosJugados)
		case "Puntos":
			suma += float64(jugador.puntos)
		case "Asistencias":
			suma += float64(jugador.asistencias)
		case "Rebotes":
			suma += float64(jugador.rebotes)
		case "Tapones":
			suma += float64(jugador.tapones)
		case "Robos":
			suma += float64(jugador.robos)
		case "Perdidas":
			suma += float64(jugador.perdidas)
		}
	}

	logger.Debug().Msg("Media de " + estadistica + " en la epoca " + strconv.Itoa(epoca.fechaInicio) + " es " + strconv.FormatFloat(suma/float64(len(epoca.estadisticasJugadores)), 'f', 6, 64))
	return suma / float64(len(epoca.estadisticasJugadores))
}

func calculaMediaJugadorEpoca(epoca Epoca, jugador Clave, estadistica string) (float64, error) {

	logger.Info().Msg("Calculando media de " + estadistica + " de " + jugador.nombreApellidos + " en la epoca " + strconv.Itoa(jugador.temporada))
	var total float64
	var totalPartidos float64

	switch estadistica {
	case "Partidos":
		for _, jugadorEpoca := range epoca.estadisticasJugadores {
			if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
				total += float64(jugadorEpoca.partidosJugados)
			}
		}
	case "Puntos":
		for _, jugadorEpoca := range epoca.estadisticasJugadores {
			if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
				total += float64(jugadorEpoca.puntos)
			}
		}
	case "Asistencias":
		for _, jugadorEpoca := range epoca.estadisticasJugadores {
			if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
				total += float64(jugadorEpoca.asistencias)
			}
		}
	case "Rebotes":
		for _, jugadorEpoca := range epoca.estadisticasJugadores {
			if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
				total += float64(jugadorEpoca.rebotes)
			}
		}
	case "Tapones":
		for _, jugadorEpoca := range epoca.estadisticasJugadores {
			if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
				total += float64(jugadorEpoca.tapones)
			}
		}
	case "Robos":
		for _, jugadorEpoca := range epoca.estadisticasJugadores {
			if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
				total += float64(jugadorEpoca.robos)
			}
		}
	case "Perdidas":
		for _, jugadorEpoca := range epoca.estadisticasJugadores {
			if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
				total += float64(jugadorEpoca.perdidas)
			}
		}
	}

	for _, jugadorEpoca := range epoca.estadisticasJugadores {
		if jugadorEpoca.nombreApellidos == jugador.nombreApellidos {
			totalPartidos += float64(jugadorEpoca.partidosJugados)
		}
	}

	logger.Info().Msg("Media de " + estadistica + " de " + jugador.nombreApellidos + " en la epoca " + strconv.Itoa(jugador.temporada) + " es " + strconv.FormatFloat(total/totalPartidos, 'f', 6, 64))
	return total / totalPartidos, nil

}

func normalizaJugador(jugador EstadisticasJugador, porcentaje float64, estadistica string) EstadisticasJugador {

	logger.Info().Msg("Normalizando estadistica " + estadistica + " de " + jugador.nombreApellidos)
	switch estadistica {
	case "Partidos":
		jugador.partidosJugados = int(float64(jugador.partidosJugados) * porcentaje)
		logger.Debug().Msg("Partidos normalizados a " + strconv.Itoa(jugador.partidosJugados))
	case "Puntos":
		jugador.puntos = int(float64(jugador.puntos) * porcentaje)
		logger.Debug().Msg("Puntos normalizados a " + strconv.Itoa(jugador.puntos))
	case "Asistencias":
		jugador.asistencias = int(float64(jugador.asistencias) * porcentaje)
		logger.Debug().Msg("Asistencias normalizadas a " + strconv.Itoa(jugador.asistencias))
	case "Rebotes":
		jugador.rebotes = int(float64(jugador.rebotes) * porcentaje)
		logger.Debug().Msg("Rebotes normalizados a " + strconv.Itoa(jugador.rebotes))
	case "Tapones":
		jugador.tapones = int(float64(jugador.tapones) * porcentaje)
		logger.Debug().Msg("Tapones normalizados a " + strconv.Itoa(jugador.tapones))
	case "Robos":
		jugador.robos = int(float64(jugador.robos) * porcentaje)
		logger.Debug().Msg("Robos normalizados a " + strconv.Itoa(jugador.robos))
	case "Perdidas":
		jugador.perdidas = int(float64(jugador.perdidas) * porcentaje)
		logger.Debug().Msg("Perdidas normalizadas a " + strconv.Itoa(jugador.perdidas))
	}

	return jugador
}

func normalizaEpoca(epocaFija Epoca, epocaNormalizar Epoca, estadistica string) (Epoca, error) {

	logger.Info().Msg("Normalizando epoca " + strconv.Itoa(epocaNormalizar.fechaInicio) + " respecto a la epoca " + strconv.Itoa(epocaFija.fechaInicio))
	mediaEpocaFija := calculaMediaEstadistica(epocaFija, estadistica)
	mediaEpocaNormalizar := calculaMediaEstadistica(epocaNormalizar, estadistica)

	if mediaEpocaFija == 0 || mediaEpocaNormalizar == 0 {
		return epocaNormalizar, errors.New("la media es 0")
	}

	porcentaje := 1.0
	if mediaEpocaFija != mediaEpocaNormalizar {
		porcentaje = mediaEpocaFija / mediaEpocaNormalizar
	}

	nuevaEpoca := Epoca{
		fechaInicio:           epocaNormalizar.fechaInicio,
		fechaFin:              epocaNormalizar.fechaFin,
		estadisticasJugadores: epocaNormalizar.estadisticasJugadores,
	}

	for clave, jugador := range nuevaEpoca.estadisticasJugadores {
		nuevaEpoca.estadisticasJugadores[clave] = normalizaJugador(jugador, porcentaje, estadistica)
	}

	return nuevaEpoca, nil
}

func porcentajeDiff(a, b int) (float64, error) {
	if b == 0 {
		return float64(a), errors.New("división por cero")
	}
	return math.Abs(float64(a-b)) / float64(a) * 100.0, nil
}

func estadisticaSimilares(jugador1 float64, jugador2 float64, umbralPorcentaje float64) (bool, error) {
	variacion, err := porcentajeDiff(int(jugador1), int(jugador2))
	if err != nil {
		logger.Error().Msg("Error al calcular la variación de las estadísticas:" + err.Error())
		return false, err
	}
	return variacion <= umbralPorcentaje, nil
}

func comparaJugadores(epocaFija Epoca, epocaNormalizada Epoca, estadistica string, umbral float64) (map[Clave][]Clave, error) {

	logger.Info().Msg("Comparando jugadores de la epoca " + strconv.Itoa(epocaNormalizada.fechaInicio) + " respecto a la epoca " + strconv.Itoa(epocaFija.fechaInicio))
	epocaNormalizada, err := normalizaEpoca(epocaFija, epocaNormalizada, estadistica)
	if err != nil {
		logger.Error().Msg("Error al normalizar las estadísticas:" + err.Error())
		return nil, err
	}
	comparador := make(map[Clave][]Clave)

	for _, jugadorNormalizado := range epocaNormalizada.estadisticasJugadores {
		claveNormalizado := Clave{
			nombreApellidos: jugadorNormalizado.nombreApellidos,
			temporada:       jugadorNormalizado.temporada,
		}
		mediaJugadorNormalizado, err := calculaMediaJugadorEpoca(epocaNormalizada, claveNormalizado, estadistica)
		if err != nil {
			logger.Error().Msg("Error al calcular la media de las estadísticas:" + err.Error())
			return nil, err
		}
		similares := make([]Clave, 0)
		comparador[claveNormalizado] = similares

		for _, jugadorFijo := range epocaFija.estadisticasJugadores {
			logger.Info().Msg("Comparando jugador " + jugadorNormalizado.nombreApellidos + " con " + jugadorFijo.nombreApellidos)
			claveFijo := Clave{
				nombreApellidos: jugadorFijo.nombreApellidos,
				temporada:       jugadorFijo.temporada,
			}
			mediaJugadorFijo, err := calculaMediaJugadorEpoca(epocaFija, claveFijo, estadistica)
			if err != nil {
				logger.Error().Msg("Error al calcular la media de las estadísticas:" + err.Error())
				return nil, err
			}
			if ok, err := estadisticaSimilares(mediaJugadorFijo, mediaJugadorNormalizado, umbral); ok {
				if err != nil {
					logger.Error().Msg("Error al calcular la variación de las estadísticas:" + err.Error())
					return nil, err
				}
				logger.Debug().Msg("Jugador " + jugadorNormalizado.nombreApellidos + " es similar a " + jugadorFijo.nombreApellidos)
				similares = append(similares, claveFijo)
				comparador[claveNormalizado] = similares
			}
		}
	}

	return comparador, nil
}
