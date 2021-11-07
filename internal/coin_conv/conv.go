package coin_conv

import "errors"

type CoinConv struct {
	Market Market
}

type Market interface {
	GetCryptoPrice(crypt, fiat string) (float64, error)
}

var errIncorrectPrice = errors.New("got incorrect price from market place")

func NewCoinConv(market Market) CoinConv {
	return CoinConv{
		market,
	}
}

func (c CoinConv) ConvertData(count float64, cryptCur, fiatCur string) (float64, error) {
	price, err := c.Market.GetCryptoPrice(cryptCur, fiatCur)
	if err != nil {
		return 0, err
	}
	if price == 0 {
		return 0, errIncorrectPrice
	}
	return count / price, nil
}
