package coin_market

import (
	"errors"
	"net/url"

	"github.com/buger/jsonparser"
)

type CoinMarket struct {
	coinClient *CoinClient
}

func NewCoinMarket() (*CoinMarket, error) {
	coinClient, err := NewCryptoPriceClient()
	if err != nil {
		return nil, err
	}
	return &CoinMarket{
		coinClient: coinClient,
	}, nil
}

func (c *CoinMarket) GetCryptoPrice(crypt, fiat string) (float64, error) {
	coinClient, err := NewCryptoPriceClient()
	if err != nil {
		return 0, err
	}
	q := url.Values{}
	q.Add("symbol", crypt)
	q.Add("convert", fiat)
	coinClient.SetQueryValues(q)

	response, err := coinClient.Do()
	if err != nil {
		return 0, err
	}
	errMsg, dataType, _, err := jsonparser.Get(response, "status", "error_message")
	switch {
	case err != nil:
		return 0, err
	case dataType != jsonparser.Null:
		return 0, errors.New(string(errMsg))
	}

	price, err := jsonparser.GetFloat(response, "data", crypt, "quote", fiat, "price")
	if err != nil {
		return 0, err
	}

	return price, nil
}
