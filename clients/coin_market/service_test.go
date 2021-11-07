package coin_market

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

type ClientMock struct {
	body             []byte
	err              error
	setQueryCount    int
	addRequestHeader int
	doCall           int
}

func (c *ClientMock) SetQueryValues(q url.Values) {
	c.setQueryCount++
}

func (c *ClientMock) AddRequestHeader(key, value string) {
	c.addRequestHeader++
}

func (c *ClientMock) Do() ([]byte, error) {
	c.doCall++
	return c.body, c.err
}

func TestCoinMarket_GetCryptoPrice(t *testing.T) {
	cases := []struct {
		caseName   string
		mockClient *ClientMock
		err        error
		exp        float64
	}{
		{
			"positive",
			&ClientMock{},
			nil,
			0,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.caseName, func(t *testing.T) {
			market := CoinMarket{
				fnCryptoPriceClient: func() (Client, error) {
					return testCase.mockClient, testCase.err
				},
			}
			v, err := market.GetCryptoPrice("", "")
			assert.Equal(t, testCase.exp, v)
			assert.Equal(t, testCase.err, err)
		})
	}
}
