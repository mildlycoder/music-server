package server

import (
	"github.com/gofiber/fiber/v2"

	"seho-music-ser/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "seho-music-ser",
			AppName:      "seho-music-ser",
		}),

		db: database.New(),
	}

	return server
}
