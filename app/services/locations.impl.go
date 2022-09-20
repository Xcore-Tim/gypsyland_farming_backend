package services

import (
	"context"
	"errors"
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationServiceImpl struct {
	locationCollection *mongo.Collection
	ctx                context.Context
}

func NewLocationService(locationCollection *mongo.Collection, ctx context.Context) LocationService {
	return &LocationServiceImpl{
		locationCollection: locationCollection,
		ctx:                ctx,
	}
}

func (l *LocationServiceImpl) CreateLocation(location *models.Location) error {

	_, err := l.locationCollection.InsertOne(l.ctx, location)
	return err

}

func (l *LocationServiceImpl) GetLocation(name *primitive.ObjectID) (*models.Location, error) {

	var location *models.Location
	query := bson.D{bson.E{Key: "_id", Value: name}}

	err := l.locationCollection.FindOne(l.ctx, query).Decode(&location)

	return location, err

}

func (l LocationServiceImpl) GetAll() ([]*models.Location, error) {

	var locations []*models.Location

	cursor, err := l.locationCollection.Find(l.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(l.ctx) {
		var location models.Location
		err := cursor.Decode(&location)

		if err != nil {
			return nil, err
		}
		locations = append(locations, &location)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(l.ctx)

	if len(locations) == 0 {
		return nil, errors.New("documents not found")
	}

	return locations, err
}

func (l LocationServiceImpl) UpdateLocation(location *models.Location) error {

	filter := bson.D{bson.E{Key: "name", Value: location.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: location.Name}, bson.E{Key: "iso", Value: location.ISO}}}}

	result, _ := l.locationCollection.UpdateOne(l.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (l LocationServiceImpl) DeleteLocation(name *string) error {

	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := l.locationCollection.DeleteOne(l.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no matched documents found for delete")
	}

	return nil
}
