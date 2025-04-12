package presentation

import (
	"context"
	"log"
	"net/http"
	"react_go_tutorial/server/domain"
	"react_go_tutorial/server/useCase"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	useCase *useCase.TodoUseCase
}

func NewTodoHandler(useCase *useCase.TodoUseCase) *TodoHandler {
	return &TodoHandler{useCase: useCase}
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	completed := c.Query("completed")
	var filter *domain.TodoFilter

	if completed != "" {
		completedBool := completed == "true"
		filter = &domain.TodoFilter{Completed: &completedBool}
	}
	
	todos, err := h.useCase.GetTodos(context.Background(), filter)
	if err != nil {
		log.Printf("Error fetching todos: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todos)
}

func (h *TodoHandler) GetDoneTodos(c *fiber.Ctx) error {
	completed := true
	filter := &domain.TodoFilter{Completed: &completed}

	todos, err := h.useCase.GetTodos(context.Background(), filter)
	if err != nil {
			log.Printf("Error fetching todos: %v", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	if len(todos) == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "There are no completed todos"})
	}

	return c.JSON(todos)
}

func (h *TodoHandler) GetNotDoneTodos(c *fiber.Ctx) error {
	completed := false
	filter := &domain.TodoFilter{Completed: &completed}

	todos, err := h.useCase.GetTodos(context.Background(), filter)
	if err != nil {
			log.Printf("Error fetching todos: %v", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	if len(todos) == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "There are no not completed todos"})
	}

	return c.JSON(todos)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err !=nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if todo.Body == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	if err := h.useCase.CreateTodo(context.Background(), todo); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(todo)
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.useCase.UpdateTodo(context.Background(), objectID, true); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": true})
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.useCase.DeleteTodo(context.Background(), objectID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": true})
}