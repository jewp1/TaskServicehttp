package api

import (
	"ApiService/internal/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	TaskHandler *http.TaskHandler
}

func NewRouter(r *Router, token string) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET, POST, PUT, DELETE",
		AllowHeaders:  "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-ID",
		ExposeHeaders: "Link",
		MaxAge:        300,
	}))
	app.Get("/taskID/:id", r.TaskHandler.GetTaskById)
	app.Post("/task", r.TaskHandler.CreateTask)
	app.Put("/task/:id", r.TaskHandler.UpdateTask)
	app.Delete("/task/:id", r.TaskHandler.DeleteTask)
	return app
}
