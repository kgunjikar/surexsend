package userManager

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents a user in the database.
type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

// Config represents the structure of the configuration JSON file.
type Config struct {
	Mongo struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
	} `json:"mongo"`
}

// ReadConfig reads the configuration from a JSON file.
func ReadConfig(filePath string) (*Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// ConnectMongo creates and connects a client to the MongoDB server.
func ConnectMongo(cfg *Config) *mongo.Client {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.Mongo.Username,
		cfg.Mongo.Password,
		cfg.Mongo.Host,
		cfg.Mongo.Port)

	// Connect directly with the MongoDB URI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

// CreateUser inserts a new user into the database.
func CreateUser(client *mongo.Client, user User) {
	collection := client.Database("your_db").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created successfully")
}

// ListUsers retrieves all users from the database.
func ListUsers(client *mongo.Client) []User {
	var users []User
	collection := client.Database("your_db").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users
}

// UpdateUser updates an existing user in the database.
func UpdateUser(client *mongo.Client, userID string, updatedUser User) {
	collection := client.Database("your_db").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"_id": userID}
	update := bson.M{"$set": updatedUser}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User updated successfully")
}

// DeleteUser removes a user from the database.
func DeleteUser(client *mongo.Client, userID string) {
	collection := client.Database("your_db").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"_id": userID}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully")
}
