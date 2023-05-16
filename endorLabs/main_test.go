package main_test

import (
	"context"
	"endorLabs/database"
	"endorLabs/models"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestObjectDB(t *testing.T) {
	// Create a Redis client for testing
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1, // Use a separate database for testing
	})

	// Create a RedisObjectDB instance for testing
	redisDB := database.NewRedisObjectDB(redisClient)

	// Create a person object for testing
	person := &models.Person{
		Name:      "John",
		LastName:  "Doe",
		Birthday:  "1990-01-01",
		BirthDate: time.Now(),
	}

	// Test storing and retrieving a person object
	err := redisDB.Store(context.Background(), person)
	if err != nil {
		t.Errorf("Failed to store person object: %v", err)
	}

	retrievedPerson, err := redisDB.GetObjectByID(context.Background(), person.GetID())
	if err != nil {
		t.Errorf("Failed to retrieve person object: %v", err)
	}

	if retrievedPerson.GetID() != person.GetID() {
		t.Errorf("Retrieved person ID doesn't match")
	}

	// Test deleting a person object
	err = redisDB.DeleteObject(context.Background(), person.GetID())
	if err != nil {
		t.Errorf("Failed to delete person object: %v", err)
	}

	// Verify that the person object is deleted
	_, err = redisDB.GetObjectByID(context.Background(), person.GetID())
	if err == nil {
		t.Errorf("Person object still exists after deletion")
	}

	// Create an animal object for testing
	animal := &models.Animal{
		Name:    "Spot",
		Type:    "Dog",
		OwnerID: "owner123",
	}

	// Test storing and retrieving an animal object
	err = redisDB.Store(context.Background(), animal)
	if err != nil {
		t.Errorf("Failed to store animal object: %v", err)
	}

	retrievedAnimal, err := redisDB.GetObjectByID(context.Background(), animal.GetID())
	if err != nil {
		t.Errorf("Failed to retrieve animal object: %v", err)
	}

	if retrievedAnimal.GetID() != animal.GetID() {
		t.Errorf("Retrieved animal ID doesn't match")
	}

	// Test deleting an animal object
	err = redisDB.DeleteObject(context.Background(), animal.GetID())
	if err != nil {
		t.Errorf("Failed to delete animal object: %v", err)
	}

	// Verify that the animal object is deleted
	_, err = redisDB.GetObjectByID(context.Background(), animal.GetID())
	if err == nil {
		t.Errorf("Animal object still exists after deletion")
	}
}
