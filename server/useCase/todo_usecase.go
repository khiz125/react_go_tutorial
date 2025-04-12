package useCase

import (
	"context"
	"react_go_tutorial/server/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoUseCase struct {
	repo domain.TodoRepository
}

func NewTodoUseCase(repo domain.TodoRepository) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (uc *TodoUseCase) GetTodos(ctx context.Context, filter *domain.TodoFilter) ([]domain.Todo, error) {
	return uc.repo.FindAll(ctx, filter)
}

func (uc *TodoUseCase) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	return uc.repo.Insert(ctx, todo)
}

func (uc *TodoUseCase) UpdateTodo(ctx context.Context, id primitive.ObjectID, completed bool) error {
	return uc.repo.Update(ctx, id, completed)
}

func (uc *TodoUseCase) DeleteTodo(ctx context.Context, id primitive.ObjectID) error {
	return uc.repo.Delete(ctx, id)
}