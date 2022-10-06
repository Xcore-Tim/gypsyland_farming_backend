package filters

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuyerRequestFilter(requestBody *models.GetRequestBody) bson.D {

	filter := bson.D{
		bson.E{Key: "$and", Value: bson.A{
			bson.D{
				bson.E{Key: "buyer.id", Value: requestBody.UserData.UserID},
				bson.E{Key: "status", Value: requestBody.Status},
				bson.E{Key: "dateCreated", Value: bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
				bson.E{Key: "dateCreated", Value: bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
			}},
		},
	}

	return filter
}

func FarmerRequestFilter(requestBody *models.GetRequestBody, teamAccess models.TeamAccess) bson.D {

	var filter primitive.D

	switch requestBody.Status {
	case 0:

		filter = bson.D{
			bson.E{Key: "$and", Value: bson.A{
				bson.D{
					bson.E{Key: "team.number", Value: bson.D{{Key: "$in", Value: teamAccess.Teams}}},
					bson.E{Key: "status", Value: requestBody.Status},
					bson.E{Key: "dateCreated", Value: bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
					bson.E{Key: "dateCreated", Value: bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
				},
			}},
		}
	default:

		filter = bson.D{
			bson.E{Key: "$and", Value: bson.A{
				bson.D{
					bson.E{Key: "farmer.id", Value: requestBody.UserData.UserID},
					bson.E{Key: "team.number", Value: bson.D{{Key: "$in", Value: teamAccess.Teams}}},
					bson.E{Key: "status", Value: requestBody.Status},
					bson.E{Key: "dateCreated", Value: bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
					bson.E{Key: "dateCreated", Value: bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
				},
			}},
		}
	}

	return filter
}

func TeamleadRequestFilter(requestBody *models.GetRequestBody) bson.D {

	filter := bson.D{
		bson.E{Key: "$and", Value: bson.A{
			bson.D{
				bson.E{Key: "status", Value: requestBody.Status},
				bson.E{Key: "team.id", Value: requestBody.UserData.TeamID},
				bson.E{Key: "dateCreated", Value: bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
				bson.E{Key: "dateCreated", Value: bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
			},
		},
		},
	}

	return filter
}

func AggregateFarmersData(requestBody *models.GetRequestBody) (bson.D, bson.D) {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{
		bson.E{Key: "status", Value: models.Complete},
		bson.E{Key: "$and", Value: bson.A{
			bson.M{"dateCreated": bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
			bson.M{"dateCreated": bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
		}},
	}}}

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
				{Key: "$sum", Value: "$accountRequest.quantity"},
			}},
		}}}

	return matchStage, groupStage
}

func AggregateBuyersData(requestBody *models.GetRequestBody, teamleadID int) (bson.D, bson.D) {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{
		bson.E{Key: "status", Value: models.Complete},
		bson.E{Key: "team.teamlead.id", Value: teamleadID},
		bson.E{Key: "$and", Value: bson.A{
			bson.M{"dateCreated": bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
			bson.M{"dateCreated": bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
		}},
	}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$buyer"},
			{Key: "totalSum", Value: bson.D{
				{Key: "$sum", Value: "$totalSum"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$accountRequest.quantity"},
			}},
			{Key: "team", Value: bson.D{
				{Key: "$first", Value: "$team"},
			}},
		}}}

	return matchStage, groupStage
}

func AggregateTeamsData(requestBody *models.GetRequestBody) (bson.D, bson.D) {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{
		bson.E{Key: "status", Value: models.Complete},
		bson.E{Key: "$and", Value: bson.A{
			bson.M{"dateCreated": bson.M{"$gte": requestBody.Period.StartDate.Unix()}},
			bson.M{"dateCreated": bson.M{"$lte": requestBody.Period.EndDate.Unix()}},
		}},
	}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$team"},
			{Key: "totalSum", Value: bson.D{
				{Key: "$sum", Value: "$totalSum"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$accountRequest.quantity"},
			}},
		}}}

	return matchStage, groupStage
}

func TeamDropdown(teamAccess *models.TeamAccess) bson.D {

	filter := bson.D{
		bson.E{Key: "$and", Value: bson.A{
			bson.D{
				bson.E{Key: "team.number", Value: bson.D{{Key: "$nin", Value: teamAccess.Teams}}},
			},
		}},
	}

	return filter
}
