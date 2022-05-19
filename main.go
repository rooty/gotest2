package main

import (
	"fmt"

	"github.com/rooty/gotest2/app"
	"github.com/rooty/gotest2/utils"
	"github.com/rs/zerolog/log"
)

func main() {
	app := app.SetupApp()
	addr := utils.GetEnvVar("GIN_ADDR")
	port := utils.GetEnvVar("GIN_PORT")

	log.Info().Msgf("Starting service on https//:%s:%s", addr, port)
	app.Run(fmt.Sprintf("%s:%s", addr, port))
}
