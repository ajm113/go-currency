package currency

import (
	"fmt"
	"strconv"
)

type (
	Money struct {
		CurrencyAttributes *Attributes
		Amount             int64
	}
)

func New(currency *Attributes) *Money {
	return &Money{
		CurrencyAttributes: currency,
	}
}

func (m *Money) SetFloat64(f float64) {
	m.Amount = int64((f * float64(m.CurrencyAttributes.TickSize)) + float64(m.CurrencyAttributes.RoundUpAmount))
}

func (m *Money) SetString(floatString string) error {
	f, err := strconv.ParseFloat(floatString, 64)

	if err != nil {
		return fmt.Errorf("go-currency: failed parsing float string: %s", err)
	}

	m.SetFloat64(f)
	return nil
}

func (m *Money) MultiplyFloat64(f float64) {
	m.Amount = int64((float64(m.Amount) * f) + m.CurrencyAttributes.RoundUpAmount)
}

func (m *Money) DivideFloat64(f float64) {
	m.Amount = int64((float64(m.Amount) / f) + m.CurrencyAttributes.RoundUpAmount)
}

func (m *Money) Float64() float64 {
	return float64(m.Amount) / float64(m.CurrencyAttributes.TickSize)
}

func (m *Money) StringNoCurrencySign() string {
	x := float64(m.Amount) / float64(m.CurrencyAttributes.TickSize)

	str := strconv.FormatFloat(x, 'b', m.CurrencyAttributes.DecimalLength, 64)

	return str
}

func (m *Money) String() string {
	return fmt.Sprintf("%s%s", m.CurrencyAttributes.Sign, m.StringNoCurrencySign())
}

func (m *Money) IsZero() bool {
	return m.Amount == 0
}
