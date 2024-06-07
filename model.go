package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id" `
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
