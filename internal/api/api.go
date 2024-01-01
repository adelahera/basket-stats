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
	app.Get("/estadisticas/:temporada/:nombre", getEstadisticas)

	// Actualiza las estadisticas de un jugador
	app.Put("/estadisticas/:temporada/:nombre", updateEstadisticas)

	// Borra las estadísticas de un jugador en una temporada
	app.Delete("/estadisticas/:temporada/:nombre", deleteEstadisticas)

	// Devuelve los jugadores similares de una epoca a otro en base a una estadística y un umbral de similitud
	app.Get("/jugador/:temporada:/:nombre/:estadistica/:similitud/:epoca", getSimilares)

	// Borra un jugador de una temporada concreta
	app.Delete("/jugador/:temporada/:nombre", deleteJugadorTemporada)

	// Borra todos los datos en todas las temporadas de un jugador
	app.Delete("/jugador/:nombre", deleteJugador)

	// Crea un nuevo jugador
	app.Post("/jugador", addJugador)

	// Actualiza los datos de un jugador
	app.Put("/jugador/:temporada/:nombre", updateJugador)

}

func getEstadisticas(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func getSimilares(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func deleteEstadisticas(c *fiber.Ctx) error {
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
