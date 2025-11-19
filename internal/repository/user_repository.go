package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/NavaneethWKT/CRUD-Go-lang/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://navaneeth_db:navaneeth@navaneethn.7dnzveh.mongodb.net/?appName=NavaneethN"
const dbName = "user_management"
const colName = "users"

var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully")
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: client.Database(dbName).Collection(colName),
	}
}

// insert a new user
func (ur *UserRepository) InsertUser(user model.User) error {
	_, err := ur.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}