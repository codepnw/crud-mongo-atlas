package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	filter := bson.M{"year": 1894}
	singleResult := collection.FindOne(ctx, filter)

	raw, err := singleResult.Raw()
	if err != nil {
		log.Fatal("main, findone: ", err)
	}
	fmt.Println(raw)

	result := Movie{}
	err = singleResult.Decode(&result)
	if err != nil {
		log.Fatal("main, decode: ", err)
	}
	fmt.Println("---------------")

	resultByte, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(resultByte))

	client.Ping(ctx, readpref.Primary()) // Ping Primary Server
	fmt.Println("server running....")
}

type Movie struct {
	ID          primitive.ObjectID `json:"id" bson:"_id" `
	ImdbID      int                `json:"imdbID" bson:"imdbID"`
	Title       string             `json:"title" bson:"title"`
	Year        int                `json:"year" bson:"year"`
	Rating      string             `json:"rating" bson:"rating"`
	Runtime     string             `json:"runtime" bson:"runtime"`
	Genre       string             `json:"genre" bson:"genre"`
	Released    string             `json:"released" bson:"released"`
	Director    string             `json:"director" bson:"director"`
	Writer      string             `json:"writer" bson:"writer"`
	Cast        string             `json:"cast" bson:"cast"`
	Metacritic  string             `json:"metacritic" bson:"metacritic"`
	ImdbRating  float64            `json:"imdbRating" bson:"imdbRating"`
	ImdbVotes   int                `json:"imdbVotes" bson:"imdbVotes"`
	Poster      string             `json:"poster" bson:"poster"`
	Plot        string             `json:"plot" bson:"plot"`
	Fullplot    string             `json:"fullplot" bson:"fullplot"`
	Language    string             `json:"language" bson:"language"`
	Country     string             `json:"country" bson:"country"`
	Awards      string             `json:"awards" bson:"awards"`
	Lastupdated string             `json:"lastupdated" bson:"lastupdated"`
	Type        string             `json:"type" bson:"type"`
}
