package application

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Start(ctx context.Context) (err error) {
	server := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	ch := make(chan error, 1)

	go func() {
		port := os.Getenv("PORT")
		addr := fmt.Sprintf(":%s", port)
		log.Info().Msgf("Starting server on port %s", port)
		err = server.Listen(addr)
		if err != nil {
			log.Fatal().Msgf("Failed to start the server %v", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return
	case <-ctx.Done():
		log.Info().Msgf("Gracefully shutdown...")
		server.Shutdown()
	}

	return
}
