package filters

import (
	accounts "gypsylandFarming/app/models/accounts"

	"go.mongodb.org/mongo-driver/bson"
)

func BuyerRequestProjection(requestBody *accounts.GetRequestBody) bson.D {

	var projection bson.D

	switch requestBody.Status {
	case 0:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "farmer", Value: 1},
			bson.E{Key: "team", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateUpdated", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	case 1:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "farmer", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateUpdated", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	case 2:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "farmer", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateUpdated", Value: 1},
			bson.E{Key: "dateFinished", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "price", Value: 1},
			bson.E{Key: "valid", Value: 1},
			bson.E{Key: "totalSum", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	case 3:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "farmer", Value: 1},
			bson.E{Key: "cancelledBy", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateCancelled", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	}
	return projection
}

func FarmerRequestProjection(requestBody *accounts.GetRequestBody) bson.D {

	var projection bson.D

	switch requestBody.Status {
	case 0:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "buyer", Value: 1},
			bson.E{Key: "team", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateUpdated", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	case 1:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "buyer", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateUpdated", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	case 2:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "buyer", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateUpdated", Value: 1},
			bson.E{Key: "dateFinished", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "price", Value: 1},
			bson.E{Key: "valid", Value: 1},
			bson.E{Key: "totalSum", Value: 1},
			bson.E{Key: "team", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	case 3:
		projection = bson.D{
			bson.E{Key: "accountRequest", Value: 1},
			bson.E{Key: "buyer", Value: 1},
			bson.E{Key: "cancelledBy", Value: 1},
			bson.E{Key: "denialReason", Value: 1},
			bson.E{Key: "dateCreated", Value: 1},
			bson.E{Key: "dateCancelled", Value: 1},
			bson.E{Key: "description", Value: 1},
			bson.E{Key: "downloadLink", Value: 1},
		}
	}
	return projection
}

func BuyerAggregationProjection() bson.D {

	projection := bson.D{

		bson.E{Key: "accountRequest", Value: 1},
		bson.E{Key: "buyer", Value: 1},
		bson.E{Key: "price", Value: 1},
		bson.E{Key: "valid", Value: 1},
		bson.E{Key: "totalSum", Value: 1},
	}

	return projection
}

func TeamsProjection() bson.D {
	projection := bson.D{
		bson.E{Key: "number", Value: 1},
	}
	return projection
}

func UCProjection() bson.D {
	projection := bson.D{
		bson.E{Key: "_id", Value: 1},
		bson.E{Key: "accountRequest", Value: 1},
		bson.E{Key: "buyer", Value: 1},
		bson.E{Key: "team", Value: 1},
		bson.E{Key: "price", Value: 1},
		bson.E{Key: "totalSum", Value: 1},
		bson.E{Key: "description", Value: 1},
	}
	return projection
}
