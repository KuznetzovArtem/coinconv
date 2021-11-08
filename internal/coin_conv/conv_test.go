package coin_conv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MarketMock struct {
	value float64
	err   error
}

func (m MarketMock) GetCryptoPrice(crypt, fiat string) (float64, error) {
	return m.value, m.err
}

func TestCoinConv_ConvertData(t *testing.T) {
	testCases := []struct {
		testName string
		m        MarketMock
		count    float64
		exp      float64
		err      error
	}{
		{
			testName: "negative scenario",
			m: MarketMock{
				value: 0,
				err:   nil,
			},
			count: 10,
			exp:   0,
			err:   errIncorrectPrice,
		},
		{
			testName: "positive scenario",
			m: MarketMock{
				value: 10,
				err:   nil,
			},
			count: 100,
			exp:   10,
			err:   nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			coinConv := NewCoinConv(testCase.m)
			v, err := coinConv.ConvertData(testCase.count, "", "")
			assert.Equal(t, testCase.exp, v)
			assert.Equal(t, testCase.err, err)
		})
	}
}
