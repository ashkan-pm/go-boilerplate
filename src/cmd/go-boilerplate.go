package main

import (
	"sync"

	"aspm.tech/go-boilerplate/src/internal/helpers"
	"aspm.tech/go-boilerplate/src/pkg/healthz"
	"github.com/rs/zerolog/log"
)

func main() {
	helpers.ConfigureGlobalLogger()
	helpers.OverloadEnv()
	env, httpPort := helpers.GetEnv()

	go healthz.Init(httpPort)

	log.Info().Msg("Application started")
	log.Info().Msg("Environment: " + env)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
