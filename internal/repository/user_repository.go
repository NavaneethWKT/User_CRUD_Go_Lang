package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/NavaneethWKT/CRUD-Go-lang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// get all users
func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	cursor, err := ur.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var users []model.User
	for cursor.Next(context.TODO()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// get a specifc user by id
func (ur *UserRepository) GetUserByID(id string) (model.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, err
	}
	var user model.User
	err = ur.collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// update a specifc user by id
func (ur *UserRepository) UpdateUser(id string, user model.User) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	result, err := ur.collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.M{"$set": user})
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}