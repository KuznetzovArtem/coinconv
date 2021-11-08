package coin_market

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

type ClientMock struct {
	body   string
	err    error
	doCall int
}

func (c *ClientMock) SetQueryValues(q url.Values) {
}

func (c *ClientMock) AddRequestHeader(key, value string) {
}

func (c *ClientMock) Do() ([]byte, error) {
	c.doCall++
	return []byte(c.body), c.err
}

func TestCoinMarket_GetCryptoPrice(t *testing.T) {
	cases := []struct {
		callDoCount int
		caseName    string
		crypt       string
		fiat        string
		mockClient  *ClientMock
		err         error
		exp         float64
	}{
		{
			1,
			"positive",
			"BTC",
			"USD",
			&ClientMock{
				body: `{"status":{"timestamp":"2021-11-07T18:26:57.045Z","error_code":0,"error_message":null,"elapsed":24,"credit_count":1,"notice":null},"data":{"BTC":{"id":1,"name":"Bitcoin","symbol":"BTC","slug":"bitcoin","num_market_pairs":8287,"date_added":"2013-04-28T00:00:00.000Z","tags":["mineable","pow","sha-256","store-of-value","state-channels","coinbase-ventures-portfolio","three-arrows-capital-portfolio","polychain-capital-portfolio","binance-labs-portfolio","arrington-xrp-capital","blockchain-capital-portfolio","boostvc-portfolio","cms-holdings-portfolio","dcg-portfolio","dragonfly-capital-portfolio","electric-capital-portfolio","fabric-ventures-portfolio","framework-ventures","galaxy-digital-portfolio","huobi-capital","alameda-research-portfolio","a16z-portfolio","1confirmation-portfolio","winklevoss-capital","usv-portfolio","placeholder-ventures-portfolio","pantera-capital-portfolio","multicoin-capital-portfolio","paradigm-xzy-screener"],"max_supply":21000000,"circulating_supply":18866625,"total_supply":18866625,"is_active":1,"platform":null,"cmc_rank":1,"is_fiat":0,"last_updated":"2021-11-07T18:26:02.000Z","quote":{"USD":{"price":62284.59175999739,"volume_24h":25124838813.622734,"volume_change_24h":-12.6921,"percent_change_1h":0.15167106,"percent_change_24h":2.34221431,"percent_change_7d":2.71412774,"percent_change_30d":14.21749045,"percent_change_60d":34.37775402,"percent_change_90d":35.44341384,"market_cap":1175100036013.9607,"market_cap_dominance":42.5094,"fully_diluted_market_cap":1307976426959.95,"last_updated":"2021-11-07T18:26:02.000Z"}}}}}`,
				err:  nil,
			},
			nil,
			62284.59175999739,
		},
		{
			1,
			"negative",
			"BTC",
			"USD",
			&ClientMock{
				body: `{"status":{"timestamp":"2021-11-07T18:26:57.045Z","error_code":0,"error_message":"some trash happened","elapsed":24,"credit_count":1,"notice":null},"data":{"BTC":{"id":1,"name":"Bitcoin","symbol":"BTC","slug":"bitcoin","num_market_pairs":8287,"date_added":"2013-04-28T00:00:00.000Z","tags":["mineable","pow","sha-256","store-of-value","state-channels","coinbase-ventures-portfolio","three-arrows-capital-portfolio","polychain-capital-portfolio","binance-labs-portfolio","arrington-xrp-capital","blockchain-capital-portfolio","boostvc-portfolio","cms-holdings-portfolio","dcg-portfolio","dragonfly-capital-portfolio","electric-capital-portfolio","fabric-ventures-portfolio","framework-ventures","galaxy-digital-portfolio","huobi-capital","alameda-research-portfolio","a16z-portfolio","1confirmation-portfolio","winklevoss-capital","usv-portfolio","placeholder-ventures-portfolio","pantera-capital-portfolio","multicoin-capital-portfolio","paradigm-xzy-screener"],"max_supply":21000000,"circulating_supply":18866625,"total_supply":18866625,"is_active":1,"platform":null,"cmc_rank":1,"is_fiat":0,"last_updated":"2021-11-07T18:26:02.000Z","quote":{"USD":{"price":62284.59175999739,"volume_24h":25124838813.622734,"volume_change_24h":-12.6921,"percent_change_1h":0.15167106,"percent_change_24h":2.34221431,"percent_change_7d":2.71412774,"percent_change_30d":14.21749045,"percent_change_60d":34.37775402,"percent_change_90d":35.44341384,"market_cap":1175100036013.9607,"market_cap_dominance":42.5094,"fully_diluted_market_cap":1307976426959.95,"last_updated":"2021-11-07T18:26:02.000Z"}}}}}`,
				err:  nil,
			},
			errors.New("some trash happened"),
			0,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.caseName, func(t *testing.T) {
			market := CoinMarket{
				fnCryptoPriceClient: func() (Client, error) {
					return testCase.mockClient, testCase.mockClient.err
				},
			}
			v, err := market.GetCryptoPrice(testCase.crypt, testCase.fiat)
			assert.Equal(t, testCase.exp, v)
			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.callDoCount, testCase.mockClient.doCall)
		})
	}
}
