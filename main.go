package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofor-little/env"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}
	mongoURI := env.Get("MONGO_URI", "")
	dbName := env.Get("DB_NAME", "")
	dbCollection := env.Get("DB_COLLECTION", "")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.SetFlags(log.Lmicroseconds)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("main: cant connect to database: ", err)
	}

	collection := client.Database(dbName).Collection(dbCollection)

	// Read Single Document
	readSingleDoc(collection, 1894)

	// Insert Document
	insResult, err := insertDoc(collection)
	if err != nil {
		log.Fatal("main, insert failed: ", err)
	}
	fmt.Println(insResult.InsertedID)

	// Update Document
	hexId := "666322800874b2a7b469a44b"
	updResult, err := updateDoc(collection, hexId, "Test Update")
	if err != nil {
		log.Fatal("main, update document: ", err)
	}
	fmt.Printf("update successfully: %v\n", updResult.UpsertedID)

	// Delete Document
	delResult, err := deleteDoc(collection, hexId)
	if err != nil {
		log.Fatal("main, delete document: ", err)
	}
	fmt.Printf("document deleted: %v\n", delResult)

	// Ping Primary Server
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("main: cant connect database..")
	}
}
