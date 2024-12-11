package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type httpServer struct {
}

func NewAPIServer() *httpServer {
	return &httpServer{}
}

func (s *httpServer) App() *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	return app
}
