package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func insertDoc(collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, Movie{ID: primitive.NewObjectID(), ImdbID: 1234567, Title: "Good Life"})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func updateDoc(collection *mongo.Collection, hexId, title string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, fmt.Errorf("error create object id: %v", err)
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: title},
	}}}

	updResult, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, fmt.Errorf("error update collection: %v", err)
	}

	return updResult, nil
}

func deleteDoc(collection *mongo.Collection, hexId string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, fmt.Errorf("error create object id: %v", err)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, fmt.Errorf("error delete collection: %v", err)
	}

	return res, nil
}