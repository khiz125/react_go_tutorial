package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	FindAll(ctx context.Context, filter *TodoFilter) ([]Todo, error)
	Insert(ctx context.Context, todo *Todo) error
	Update(ctx context.Context, id primitive.ObjectID, completed bool) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}