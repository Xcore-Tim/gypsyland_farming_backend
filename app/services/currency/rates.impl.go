package services

import (
	"context"
	"encoding/xml"
	"fmt"
	models "gypsylandFarming/app/models/currency"
	"io"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/encoding/charmap"
)

type CurrencyRatesServiceImpl struct {
	currencyCollection *mongo.Collection
	ctx                context.Context
}

func NewCurrencyRatesService(currencyCollection *mongo.Collection, ctx context.Context) CurrencyRatesService {

	return &CurrencyRatesServiceImpl{
		currencyCollection: currencyCollection,
		ctx:                ctx,
	}
}

func (srvc CurrencyRatesServiceImpl) GetCurrencyRates(requestDate string) (map[string]float64, error) {

	var cr models.ValuteRate
	var err error

	err = srvc.RequestCurrencyRate(&cr, requestDate)

	if err != nil {
		return nil, err
	}

	valMap := make(map[string]float64)

	for _, v := range cr.ValuteList {

		strVal := strings.Replace(v.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(strVal, 64)

		if err != nil {
			continue
		}
		value = srvc.RoundFloat(value, 2)

		valMap[v.CharCode] = value

	}

	return valMap, nil
}

func (srvc CurrencyRatesServiceImpl) RequestCurrencyRate(cr *models.ValuteRate, requestDate string) error {

	baseUrl := "http://www.cbr.ru/scripts/XML_daily.asp?date_req="

	ulrPath := baseUrl + requestDate

	request, err := http.NewRequest(http.MethodGet, ulrPath, nil)

	if err != nil {
		return err
	}

	if err := srvc.ClientRequest(cr, request); err != nil {
		return err
	}

	return nil
}

func (srvc CurrencyRatesServiceImpl) ClientRequest(cr *models.ValuteRate, request *http.Request) error {

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	defer response.Body.Close()

	decoder := NewDecoderXML(response.Body)

	if err := decoder.Decode(&cr); err != nil {
		return err
	}

	return nil
}

func NewDecoderXML(body io.ReadCloser) *xml.Decoder {

	d := xml.NewDecoder(body)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	return d
}
