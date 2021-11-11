package coin_comission

import (
	"coinconv/internal"
	"errors"
)

type CommissionCalculator struct {
	converter  internal.Converter
	commission float64
}

func (c *CommissionCalculator) ConvertData(count float64, cryptCur, fiatCur string) (float64, error) {
	price, err := c.converter.ConvertData(count, cryptCur, fiatCur)
	if err != nil {
		return 0, err
	}
	return price * (1 + c.commission), nil
}

func NewCommissionCalculator(conv internal.Converter, commission float64) (internal.Converter, error) {
	if commission == 0 {
		return nil, errors.New("empty commission")
	}
	return &CommissionCalculator{
		converter:  conv,
		commission: commission,
	}, nil
}
