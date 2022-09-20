package services

import (
	"context"
	"errors"
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PositionServiceImpl struct {
	positionCollection *mongo.Collection
	ctx                context.Context
}

func NewPositionService(positionCollection *mongo.Collection, ctx context.Context) PositionService {

	return &PositionServiceImpl{
		positionCollection: positionCollection,
		ctx:                ctx,
	}

}

func (p PositionServiceImpl) CreatePosition(position *models.Position) error {
	_, err := p.positionCollection.InsertOne(p.ctx, position)
	return err
}

func (p PositionServiceImpl) GetPosition(id *primitive.ObjectID) (*models.Position, error) {

	var position *models.Position
	query := bson.D{bson.E{Key: "_id", Value: id}}

	err := p.positionCollection.FindOne(p.ctx, query).Decode(&position)

	return position, err

}

func (p PositionServiceImpl) GetAll() ([]*models.Position, error) {

	var positions []*models.Position

	cursor, err := p.positionCollection.Find(p.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(p.ctx) {
		var position models.Position
		err := cursor.Decode(&position)

		if err != nil {
			return nil, err
		}
		positions = append(positions, &position)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(p.ctx)

	if len(positions) == 0 {
		return nil, errors.New("documents not found")
	}

	return positions, err

}

func (p PositionServiceImpl) UpdatePosition(position *models.Position) error {

	filter := bson.D{bson.E{Key: "_id", Value: position.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: position.Name}, bson.E{Key: "role_number", Value: position.RoleNumber}}}}

	result, _ := p.positionCollection.UpdateOne(p.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched s found for update")
	}

	return nil
}

func (p PositionServiceImpl) DeletePosition(id *primitive.ObjectID) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	result, _ := p.positionCollection.DeleteOne(p.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no matched documents found to delete")
	}

	return nil

}
