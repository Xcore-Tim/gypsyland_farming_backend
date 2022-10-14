package services

import (
	models "gypsylandFarming/app/models/currency"
	"math"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrencyService interface {
	CreateCurrency(*models.Currency) (primitive.ObjectID, error)
	UpdateCurrency(*models.Currency) error

	GetCurrency(primitive.ObjectID) (*models.Currency, error)
	GetBaseCurrency() (*models.Currency, error)

	GetAll() ([]*models.Currency, error)

	DeleteAll() (int, error)
}

type CurrencyRatesService interface {
	GetCurrencyRates(string) (map[string]float64, error)
	RequestCurrencyRate(*models.ValuteRate, string) error

	RoundFloat(float64, uint) float64
}

func (srvc CurrencyRatesServiceImpl) RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
