package infrastructure

import (
	infrastructure_database "go-ddd/src/infrastructure/database"
	"go-ddd/src/order"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type httpServer struct {
	addr string
	log  Logger
}

func NewAPIServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	order.RegisterPackage(app, infrastructure_database.DB)

	return app.Listen(s.addr)
}
