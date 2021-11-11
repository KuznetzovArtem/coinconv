package internal

type Converter interface {
	ConvertData(count float64, cryptCur, fiatCur string) (float64, error)
}

type Client interface {
	GetCryptoPrice(crypt, fiat string) (float64, error)
}
