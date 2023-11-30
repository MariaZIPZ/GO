package business

import "fmt"

type Currency struct {
	name   string
	exRate float64
}

func CurrencyDefaultConstructor() *Currency {
	c := new(Currency)

	c.name = "USD"
	c.exRate = 37.75

	return c
}

func CurrencyConstructor(name string, exRate float64) *Currency {
	c := new(Currency)
	c.name = name
	c.exRate = exRate

	return c
}

func CurrencyConstructorWithExRate(exRate float64) *Currency {
	c := new(Currency)
	c.name = "USD"
	c.exRate = exRate

	return c
}

func (c *Currency) SetExRate(exRate float64) {
	c.exRate = exRate
}

func (c *Currency) SetName(name string) {
	c.name = name
}

func (c *Currency) GetExRate() float64 {
	return c.exRate
}

func (c *Currency) GetName() string {
	return c.name
}

func (c *Currency) GetInfo() string {
	exRateStr := fmt.Sprint(c.exRate)

	return "1 " + c.name + " = " + exRateStr + " UAH"
}
