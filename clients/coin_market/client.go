package coin_market

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	coinMarketToken = "9df3973d-c048-462d-8994-386cd966502f"
	getPriceUrl     = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"
)

type CoinClient struct {
	httpClient *http.Client
	r          *http.Request
}

type Client interface {
	AddRequestHeader(key, value string)
	SetQueryValues(q url.Values)
	Do() ([]byte, error)
}

func NewClient(method, url string) (Client, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return &CoinClient{
		httpClient: &http.Client{},
		r:          request,
	}, nil
}

func (c *CoinClient) SetQueryValues(q url.Values) {
	c.r.URL.RawQuery = q.Encode()
}

func (c *CoinClient) AddRequestHeader(key, value string) {
	c.r.Header.Add(key, value)
}

func (c *CoinClient) Do() ([]byte, error) {
	resp, err := c.httpClient.Do(c.r)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func NewCryptoPriceClient() (Client, error) {
	coinClient, err := NewClient(
		http.MethodGet,
		getPriceUrl,
	)
	if err != nil {
		return nil, err
	}
	coinClient.AddRequestHeader("Accepts", "application/json")
	coinClient.AddRequestHeader("X-CMC_PRO_API_KEY", coinMarketToken)
	return coinClient, nil
}
