package coin_conv

import (
	"coinconv/internal"
	"errors"
)

type CoinConv struct {
	Market internal.Client
}

var errIncorrectPrice = errors.New("got incorrect price from market place")

func NewCoinConv(market internal.Client) internal.Converter {
	return CoinConv{
		market,
	}
}

func (c CoinConv) ConvertData(count float64, cryptCur, fiatCur string) (float64, error) {
	if count == 0 {
		return 0, nil
	}

	price, err := c.Market.GetCryptoPrice(cryptCur, fiatCur)
	if err != nil {
		return 0, err
	}
	if price == 0 {
		return 0, errIncorrectPrice
	}
	return count / price, nil
}
