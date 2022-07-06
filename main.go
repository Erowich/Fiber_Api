package main

import (
	"Fiber_Api_/app"
	"Fiber_Api_/configs"
	"Fiber_Api_/repository"
	"Fiber_Api_/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDb := repository.NewTodoRepositoryDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Put("/api/todo/:id", td.EditTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":8080")
}
