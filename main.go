package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofor-little/env"
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

	// Call Operations
	readSingleDoc(collection, 1894)

	client.Ping(ctx, readpref.Primary()) // Ping Primary Server
	fmt.Println("server running....")
}
