package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Currency struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	ISO    string             `json:"iso" bson:"iso"`
	Symbol string             `json:"symbol" bson:"symbol"`
	Value  float64            `bson:"value,omitempty"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type ValuteRate struct {
	ValuteList []Valute `xml:"Valute"`
}
