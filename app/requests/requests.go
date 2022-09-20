package db_requests

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

func BuyerRequest(requestBody *models.GetRequestBody) bson.D {

	filter := bson.D{
		bson.E{Key: "buyerID", Value: requestBody.UserIdentity.UserID},
		bson.E{Key: "status", Value: requestBody.Status},
		bson.E{Key: "$and", Value: bson.A{
			bson.M{"date_created": bson.M{"$gte": requestBody.Period.StartDate}},
			bson.M{"date_created": bson.M{"$lte": requestBody.Period.EndDate}},
		}},
	}
	return filter
}

func FarmerRequest(requestBody *models.GetRequestBody, teamAccess models.TeamAccess) bson.D {

	var filter primitive.D

	switch requestBody.Status {
	case 0:
		filter = bson.D{
			bson.E{Key: "$and", Value: bson.A{
				bson.D{
					bson.E{Key: "date_created", Value: bson.D{{Key: "$gte", Value: requestBody.Period.StartDate.Unix()}}},
					bson.E{Key: "date_created", Value: bson.D{{Key: "$lte", Value: requestBody.Period.EndDate.Unix()}}},
					bson.E{Key: "teamID", Value: bson.D{{Key: "$in", Value: teamAccess.Teams}}},
					bson.E{Key: "status", Value: requestBody.Status},
				},
			},
			},
		}
	default:

		filter = bson.D{
			bson.E{Key: "$and", Value: bson.A{
				bson.D{
					bson.E{Key: "date_created", Value: bson.D{{Key: "$gte", Value: requestBody.Period.StartDate.Unix()}}},
					bson.E{Key: "date_created", Value: bson.D{{Key: "$lte", Value: requestBody.Period.EndDate.Unix()}}},
					bson.E{Key: "teamID", Value: bson.D{{Key: "$in", Value: teamAccess.Teams}}},
					bson.E{Key: "status", Value: requestBody.Status},
					bson.E{Key: "farmerID", Value: requestBody.UserIdentity.UserID},
				},
			},
			},
		}
	}

	return filter
}

func TeamleadRequest(requestBody *models.GetRequestBody) bson.D {

	filter := bson.D{
		bson.E{Key: "$and", Value: bson.A{
			bson.D{
				bson.E{Key: "date_created", Value: bson.D{{Key: "$gte", Value: requestBody.Period.StartDate.Unix()}}},
				bson.E{Key: "date_created", Value: bson.D{{Key: "$lte", Value: requestBody.Period.EndDate.Unix()}}},
				bson.E{Key: "status", Value: requestBody.Status},
				bson.E{Key: "teamID", Value: requestBody.UserIdentity.TeamID},
			},
		},
		},
	}

	return filter
}

func TLFAdminRequest(requestBody *models.GetRequestBody) bson.D {

	filter := bson.D{
		bson.E{Key: "$and", Value: bson.A{
			bson.D{
				bson.E{Key: "date_created", Value: bson.D{{Key: "$gte", Value: requestBody.Period.StartDate.Unix()}}},
				bson.E{Key: "date_created", Value: bson.D{{Key: "$lte", Value: requestBody.Period.EndDate.Unix()}}},
				bson.E{Key: "status", Value: requestBody.Status},
			},
		},
		},
	}

	return filter
}
