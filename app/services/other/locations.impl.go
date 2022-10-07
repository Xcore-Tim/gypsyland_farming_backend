package services

import (
	"context"
	"errors"
	"gypsylandFarming/app/models"

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

func (srvc *LocationServiceImpl) CreateLocation(location *models.Location) error {

	_, err := srvc.locationCollection.InsertOne(srvc.ctx, location)
	return err
}

func (srvc *LocationServiceImpl) GetLocation(id primitive.ObjectID) (*models.Location, error) {

	var location models.Location
	query := bson.D{bson.E{Key: "_id", Value: id}}

	if err := srvc.locationCollection.FindOne(srvc.ctx, query).Decode(&location); err != nil {
		return &location, errors.New("no locations found")
	}

	return &location, nil
}

func (srvc *LocationServiceImpl) GetLocationByName(name string) (*models.Location, error) {

	var location models.Location
	query := bson.D{bson.E{Key: "name", Value: name}}

	if err := srvc.locationCollection.FindOne(srvc.ctx, query).Decode(&location); err != nil {
		return &location, errors.New("no locations found")
	}

	return &location, nil

}

func (srvc LocationServiceImpl) GetAll() ([]*models.Location, error) {

	var locations []*models.Location

	cursor, err := srvc.locationCollection.Find(srvc.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(srvc.ctx) {
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

	cursor.Close(srvc.ctx)

	if len(locations) == 0 {
		return nil, errors.New("documents not found")
	}

	return locations, err
}

func (srvc LocationServiceImpl) UpdateLocation(location *models.Location) error {

	filter := bson.D{bson.E{Key: "name", Value: location.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: location.Name}, bson.E{Key: "iso", Value: location.ISO}}}}

	result, _ := srvc.locationCollection.UpdateOne(srvc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc LocationServiceImpl) DeleteLocation(name *string) error {

	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := srvc.locationCollection.DeleteOne(srvc.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no matched documents found for delete")
	}

	return nil
}

func (srvc LocationServiceImpl) DeleteAll() (int, error) {

	filter := bson.D{bson.E{}}

	result, err := srvc.locationCollection.DeleteMany(srvc.ctx, filter)

	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil

}
