package main

import (
	"app/config"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	// load the config
	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().Msg("Starting the app")

	// print the config
	logger.Debug().Interface("config", conf).Msg("config")
	app, clenaup, err := wireApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer clenaup()
	app.RegisterRoutes()
	if err := app.Run(); err != nil {
		panic(err)
	}

}
