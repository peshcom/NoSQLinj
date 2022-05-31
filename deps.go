package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DBInstance() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://username:password@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("Storage").Collection(collectionName)
}

//Client Database instance
var Client *mongo.Client = DBInstance()
var ctx = context.TODO()
var userCollection *mongo.Collection = OpenCollection(Client, "user")

func checkUser(username string, password string) (bool, error) {
	var filter interface{}
	err := bson.UnmarshalExtJSON([]byte(`{"username": "`+username+`","password":"`+password+`"}`),
		false, &filter)
	if err != nil {
		return false, err
	}
	var curUser user
	err = userCollection.FindOne(ctx, filter).Decode(&curUser)
	if err != nil {
		return false, err
	}
	return true, nil
}


// добавить несколько функций checkUser с разными подходами
