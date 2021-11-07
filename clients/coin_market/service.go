package coin_market

import (
	"github.com/buger/jsonparser"
	"net/url"
)

type CoinMarket struct {
	fnCryptoPriceClient func() (Client, error)
}

func NewCoinMarket() CoinMarket {
	return CoinMarket{
		fnCryptoPriceClient: NewCryptoPriceClient,
	}
}

func (c *CoinMarket) GetCryptoPrice(crypt, fiat string) (float64, error) {
	coinClient, err := c.fnCryptoPriceClient()
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

	price, err := jsonparser.GetFloat(response, "data", crypt, "quote", fiat, "price")
	if err != nil {
		return 0, err
	}

	return price, nil
}
