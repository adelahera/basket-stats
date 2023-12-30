package api

import "github.com/gofiber/fiber/v2"

var app *fiber.App

func SetupApi() *fiber.App {

	if app == nil {
		app = fiber.New()
		setRoutes()
	}

	return app
}

func setRoutes() {

	// Devuelve las estadísticas de un jugador
	app.Get("/estadisticas/:temporada/:jugador", getEstadisticas)

	// Devuelve los jugadores similares a otro en base a una estadística y un umbral de similitud
	app.Get("/similares/:jugador/:temporada:/:similitud:/:estadistica", getSimilares)

	// Borra un jugador
	app.Delete("/borrar/:temporada/:jugador", deleteJugador)

	// Crea un nuevo jugador
	app.Post("/crear/:temporada/:jugador", addJugador)

	// Actualiza los datos de un jugador
	app.Put("/actualizar/:temporada/:jugador", updateJugador)

	// Borra todos los jugadores de una época
	app.Delete("/borrar/:epoca", deleteEpoca)

}

func getEstadisticas(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func getSimilares(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func deleteJugador(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func addJugador(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func updateJugador(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func deleteEpoca(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
