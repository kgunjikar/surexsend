package userManager

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// TestConnectMongo tests the ConnectMongo function.
func TestConnectMongo(t *testing.T) {
	cfg := &Config{
		Mongo: struct {
			Username string "json:\"username\""
			Password string "json:\"password\""
			Host     string "json:\"host\""
			Port     string "json:\"port\""
			Database string "json:\"database\""
		}{
			Username: "testUser",
			Password: "testPass",
			Host:     "localhost",
			Port:     "27017",
			Database: "testDB",
		},
	}

	client := ConnectMongo(cfg)
	if client == nil {
		t.Errorf("ConnectMongo() = %v, want non-nil", client)
	}

	// Clean up after test
	client.Disconnect(context.Background())
}

// TestCreateUser tests the CreateUser function.
func TestCreateUser(t *testing.T) {
	// Assuming ConnectMongo is working correctly.
	cfg := &Config{ /* ... */ } // Your test config
	client := ConnectMongo(cfg)
	defer client.Disconnect(context.Background())

	testUser := User{ID: "testID", Name: "Test User", Email: "test@example.com", Password: "testPassword"}
	CreateUser(client, testUser)

	// Verify if user is created
	collection := client.Database("testDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result User
	err := collection.FindOne(ctx, bson.M{"_id": "testID"}).Decode(&result)
	if err != nil {
		t.Errorf("CreateUser() failed, user not found")
	}

	// Clean up after test
	collection.DeleteOne(ctx, bson.M{"_id": "testID"})
}
