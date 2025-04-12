package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"react_go_tutorial/server/infrastructure"
	"react_go_tutorial/server/presentation"
	"react_go_tutorial/server/useCase"
)

func main() {
	app := fiber.New()

	repo, err := infrastructure.NewTodoRepository()
	if err != nil {
		log.Fatal("Failed to initialize repository:", err)
	}

	uc := useCase.NewTodoUseCase(repo)

	handler := presentation.NewTodoHandler(uc)

	app.Get("/api/todos", handler.GetTodos)
	app.Get("/api/todos/done", handler.GetDoneTodos)
	app.Get("/api/todos/not-done", handler.GetNotDoneTodos)
	app.Post("/api/todos", handler.CreateTodo)
	app.Patch("/api/todos/:id", handler.UpdateTodo)
	app.Delete("/api/todos/:id", handler.DeleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
