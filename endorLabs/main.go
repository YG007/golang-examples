package main

import (
	"context"
	"endorLabs/database"
	"endorLabs/models"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	// Create a Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisObjectDB instance
	redisDB := database.NewRedisObjectDB(redisClient)

	// Store a person object
	person := &models.Person{
		Name:      "John",
		LastName:  "Doe",
		Birthday:  "1990-01-01",
		BirthDate: time.Now(),
	}
	err := redisDB.Store(context.Background(), person)
	if err != nil {
		log.Fatalf("Failed to store person object: %v", err)
	}

	// Retrieve the person object by ID
	obj, err := redisDB.GetObjectByID(context.Background(), person.GetID())
	if err != nil {
		log.Fatalf("Failed to retrieve person object: %v", err)
	}

	// Type assertion to access person object's methods
	retrievedPerson, ok := obj.(*models.Person)
	if ok {
		fmt.Println("Retrieved Person:")
		fmt.Println("ID:", retrievedPerson.GetID())
		fmt.Println("Name:", retrievedPerson.GetName())
		fmt.Println("Last Name:", retrievedPerson.LastName)
		fmt.Println("Birthday:", retrievedPerson.Birthday)
		fmt.Println("Birth Date:", retrievedPerson.BirthDate)
		fmt.Println("Kind:", retrievedPerson.GetKind())
	}

	// Retrieve the person object by name
	obj, err = redisDB.GetObjectByName(context.Background(), person.GetName())
	if err != nil {
		log.Fatalf("Failed to retrieve person object by name: %v", err)
	}

	// Type assertion to access person object's methods
	retrievedPerson, ok = obj.(*models.Person)
	if ok {
		fmt.Println("Retrieved Person by Name:")
		fmt.Println("ID:", retrievedPerson.GetID())
		fmt.Println("Name:", retrievedPerson.GetName())
		fmt.Println("Last Name:", retrievedPerson.LastName)
		fmt.Println("Birthday:", retrievedPerson.Birthday)
		fmt.Println("Birth Date:", retrievedPerson.BirthDate)
		fmt.Println("Kind:", retrievedPerson.GetKind())
	}

	// Store an animal object
	animal := &models.Animal{
		Name:    "Spot",
		Type:    "Dog",
		OwnerID: uuid.New().String(),
	}
	err = redisDB.Store(context.Background(), animal)
	if err != nil {
		log.Fatalf("Failed to store animal object: %v", err)
	}

	// Retrieve the animal object by ID
	obj, err = redisDB.GetObjectByID(context.Background(), animal.GetID())
	if err != nil {
		log.Fatalf("Failed to retrieve animal object: %v", err)
	}

	// Type assertion to access animal object's methods
	retrievedAnimal, ok := obj.(*models.Animal)
	if ok {
		fmt.Println("Retrieved Animal:")
		fmt.Println("ID:", retrievedAnimal.GetID())
		fmt.Println("Name:", retrievedAnimal.GetName())
		fmt.Println("Type:", retrievedAnimal.Type)
		fmt.Println("OwnerID:", retrievedAnimal.OwnerID)
		fmt.Println("Kind:", retrievedAnimal.GetKind())
	}

	// Retrieve the animal object by name
	obj, err = redisDB.GetObjectByName(context.Background(), animal.GetName())
	if err != nil {
		log.Fatalf("Failed to retrieve animal object by name: %v", err)
	}

	// Type assertion to access animal object's methods
	retrievedAnimal, ok = obj.(*models.Animal)
	if ok {
		fmt.Println("Retrieved Animal by Name:")
		fmt.Println("ID:", retrievedAnimal.GetID())
		fmt.Println("Name:", retrievedAnimal.GetName())
		fmt.Println("Type:", retrievedAnimal.Type)
		fmt.Println("OwnerID:", retrievedAnimal.OwnerID)
		fmt.Println("Kind:", retrievedAnimal.GetKind())
	}

	// List all person objects
	personObjects, err := redisDB.ListObjects(context.Background(), "Person")
	if err != nil {
		log.Fatalf("Failed to list person objects: %v", err)
	}

	fmt.Println("List of Person Objects:")
	for _, obj := range personObjects {
		personObj, ok := obj.(*models.Person)
		if ok {
			fmt.Println("ID:", personObj.GetID())
			fmt.Println("Name:", personObj.GetName())
			fmt.Println("Last Name:", personObj.LastName)
			fmt.Println("Birthday:", personObj.Birthday)
			fmt.Println("Birth Date:", personObj.BirthDate)
			fmt.Println("Kind:", personObj.GetKind())
			fmt.Println()
		}
	}

	// List all animal objects
	animalObjects, err := redisDB.ListObjects(context.Background(), "Animal")
	if err != nil {
		log.Fatalf("Failed to list animal objects: %v", err)
	}

	fmt.Println("List of Animal Objects:")
	for _, obj := range animalObjects {
		animalObj, ok := obj.(*models.Animal)
		if ok {
			fmt.Println("ID:", animalObj.GetID())
			fmt.Println("Name:", animalObj.GetName())
			fmt.Println("Type:", animalObj.Type)
			fmt.Println("OwnerID:", animalObj.OwnerID)
			fmt.Println("Kind:", animalObj.GetKind())
			fmt.Println()
		}
	}

	// Delete the person object by ID
	err = redisDB.DeleteObject(context.Background(), person.GetID())
	if err != nil {
		log.Fatalf("Failed to delete person object: %v", err)
	}
	fmt.Println("Person object deleted successfully.")

	// Delete the animal object by ID
	err = redisDB.DeleteObject(context.Background(), animal.GetID())
	if err != nil {
		log.Fatalf("Failed to delete animal object: %v", err)
	}
	fmt.Println("Animal object deleted successfully.")
}
