package main

import (
	"fmt"

	"github.com/adelahera/basket-stats/internal/config"
	"github.com/adelahera/basket-stats/internal/log"
)

func main() {
	config, err := config.GetConfig()
	var logger = log.GetLogger()
	if &logger == nil {
		fmt.Println("Error al obtener el logger")
		return
	}
	if err != nil {
		fmt.Println("Error al obtener la configuración:", err)
		return
	}
	println("folder:", config.LogFolder)
	fmt.Println("LogFile:", config.LogFile)

	logger.Debug().Msg("Mensaje de debug")
	logger.Info().Msg("Mensaje de información")
}
