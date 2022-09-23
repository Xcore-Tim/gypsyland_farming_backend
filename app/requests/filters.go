package filters

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuyerRequestFilter(requestBody *models.GetRequestBody) bson.D {

	filter := bson.D{
		bson.E{Key: "buyer.id", Value: requestBody.UserData.UserID},
		bson.E{Key: "status", Value: requestBody.Status},
		bson.E{Key: "$and", Value: bson.A{
			bson.M{"dateCreated": bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
			bson.M{"dateCreated": bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
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
					bson.E{Key: "team.id", Value: bson.D{{Key: "$in", Value: teamAccess.Teams}}},
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
					bson.E{Key: "team.id", Value: bson.D{{Key: "$in", Value: teamAccess.Teams}}},
					bson.E{Key: "status", Value: requestBody.Status},
					bson.E{Key: "farmer.id", Value: requestBody.UserData.UserID},
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
				bson.E{Key: "team.id", Value: requestBody.UserData.TeamID},
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

func AggregateFarmersData() (bson.D, bson.D) {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "status", Value: models.Complete}}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$farmer"},
			{Key: "price", Value: bson.D{
				{Key: "$sum", Value: "$price"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$account_request.quantity"},
			}},
		}}}

	return matchStage, groupStage
}

func AggregateBuyersData(teamlead_id int) (bson.D, bson.D) {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{
		bson.E{Key: "status", Value: models.Complete},
		bson.E{Key: "team.teamlead.id", Value: teamlead_id},
	}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$buyer"},
			{Key: "price", Value: bson.D{
				{Key: "$sum", Value: "$price"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$account_request.quantity"},
			}},
			{Key: "team", Value: bson.D{
				{Key: "$first", Value: "$team"},
			}},
		}}}

	return matchStage, groupStage
}

func AggregateTeamsData() (bson.D, bson.D) {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "status", Value: models.Complete}}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$team"},
			{Key: "price", Value: bson.D{
				{Key: "$sum", Value: "$price"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$account_request.quantity"},
			}},
		}}}

	return matchStage, groupStage
}
