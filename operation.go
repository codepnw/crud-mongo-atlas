package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func readSingleDoc(collection *mongo.Collection, year int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"year": year}
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

	resultByte, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("main: unabled to marshal json: ", resultByte)
	}
	fmt.Println(string(resultByte))
}
