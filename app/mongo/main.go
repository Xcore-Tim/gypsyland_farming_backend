package mongodb

import (
	"context"
	"encoding/json"
	mongodb_models "gypsyland_farming/app/models"

	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createNewLocation(client *mongo.Client) {

	var newLocation mongodb_models.Location

	collection := client.Database("gypsyland").Collection("locations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, newLocation)

	var response http.ResponseWriter
	json.NewEncoder(response).Encode(result)

}

func Main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	createNewLocation(client)

}
