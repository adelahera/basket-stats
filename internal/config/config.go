package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	LogFolder string `mapstructure:"log-folder"`
	LogFile   string `mapstructure:"log-file"`
}

var Conf *Config

func newConfig() (*Config, error) {
	viper.AutomaticEnv()

	pflag.String("log-folder", "./logs", "Carpeta para almacenar los archivos de log")
	pflag.String("log-file", "app.log", "Nombre del archivo de log")

	viper.BindPFlags(pflag.CommandLine)

	viper.BindEnv("log-folder", "LOG_FOLDER")
	viper.BindEnv("log-file", "LOG_FILE")

	pflag.Parse()

	logFolder := viper.GetString("log-folder")
	if err := os.MkdirAll(logFolder, os.ModePerm); err != nil {
		fmt.Println("Error al crear la carpeta de logs:", err)
		return nil, err
	}

	logFile := logFolder + "/" + viper.GetString("log-file")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error al abrir el archivo de log:", err)
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		fmt.Println("Error al unmarshal sobre la configuración:", err)
		return nil, err
	}
	return config, nil
}

func GetConfig() (*Config, error) {
	if Conf == nil {
		var err error
		Conf, err = newConfig()
		if err != nil {
			return nil, err
		}
	}
	return Conf, nil
}
