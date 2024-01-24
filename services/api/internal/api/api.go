package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

type API struct {
	App    *fiber.App
	Logger zerolog.Logger
}

func New() *API {
	return &API{
		App:    fiber.New(),
		Logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
}

func (a API) Start() {
	if err := a.App.Listen(":3000"); err != nil {
		a.Logger.Fatal().Err(err).Msg("Failed to serve API.")
	}
}

func (a API) Shutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until shutdown signal is received.
	<-c

	a.Logger.Info().Msg("Gracefully shutting down server...")
	if err := a.App.Shutdown(); err != nil {
		a.Logger.Error().Err(err).Msg("Failed shutting down the API.")
	}

	a.Logger.Info().Msg("Running cleanup tasks...")
}
