package log

import (
	"fmt"
	"os"

	"github.com/adelahera/basket-stats/internal/config"
	"github.com/rs/zerolog"
)

var logger *zerolog.Logger

func initLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	conf, _ := config.GetConfig()
	logFolder := conf.LogFolder
	logFile := conf.LogFile
	var output *os.File
	if logFolder != "" && logFile != "" {
		println("golaaaa")
		filePath := logFolder + "/" + logFile
		println("filePath:", filePath)
		var err error
		output, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			fmt.Printf("Error opening log file: %v, using stdout\n", err)
			output = os.Stdout
		}
	} else {
		fmt.Println("No log folder or log file provided, using stdout")
		output = os.Stdout
	}

	l := zerolog.New(output).With().Timestamp().Logger()
	return &l
}

func GetLogger() *zerolog.Logger {

	if logger == nil {
		logger = initLogger()
	}

	return logger
}
