package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	"react_go_tutorial/server/domain"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoRepositoryImpl struct {
	collection *mongo.Collection
}

type TodoUpdate struct {
	Completed bool `bson:"completed"`
}

func NewTodoRepository() (*TodoRepositoryImpl, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Printf("Error pinging MongoDB: %v", err)
		return nil, err
	}

	collection := client.Database("goland_db").Collection("todos")
	if collection == nil {
		log.Printf("Error pinging MongoDB: %v", err)
		return nil, fmt.Errorf("failed to get collection")
	}

	return &TodoRepositoryImpl{
		collection: collection,
	}, nil
}

func (repo *TodoRepositoryImpl) FindAll(ctx context.Context, filter *domain.TodoFilter) ([]domain.Todo, error) {
	log.Println("Starting FindAll")
	if repo.collection == nil {
		return nil, fmt.Errorf("collection is nil")
	}
	
	var todos []domain.Todo

	filterDoc := bson.M{}
	if filter != nil && filter.Completed != nil {
		filterDoc["completed"] = *filter.Completed
	}
	
	cursor, err := repo.collection.Find(ctx, filterDoc)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo domain.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *TodoRepositoryImpl) Insert(ctx context.Context, todo *domain.Todo) error {
	insertResult, err := repo.collection.InsertOne(ctx, todo)
	if err != nil {
		return err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

func (repo *TodoRepositoryImpl) Update(ctx context.Context, id primitive.ObjectID, completed bool) error {
	update := TodoUpdate{Completed: completed}
	_, err := repo.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (repo *TodoRepositoryImpl) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}