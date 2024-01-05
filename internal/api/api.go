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
	app.Get("/estadisticas/:nombre/:temporada", getEstadisticas)

	// Actualiza las estadisticas de un jugador
	app.Put("/estadisticas/:nombre/:temporada", updateEstadisticas)

	// Devuelve los jugadores similares de una epoca a otro en base a una estadística y un umbral de similitud
	app.Get("/jugador/:nombre:/:temporada/:estadistica/:similitud/:epoca", getSimilares)

	// Borra un jugador de una temporada concreta
	app.Delete("/jugador/:nombre/:temporada", deleteJugadorTemporada)

	// Borra todos los datos en todas las temporadas de un jugador
	app.Delete("/jugador/:nombre", deleteJugador)

	// Crea un nuevo jugador
	app.Post("/jugador", addJugador)

	// Actualiza los datos de un jugador
	app.Put("/jugador/:nombre/:temporada", updateJugador)

}

func getEstadisticas(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func getSimilares(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func deleteJugadorTemporada(c *fiber.Ctx) error {
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

func updateEstadisticas(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
