package business

import "fmt"

type Product struct {
	name     string
	price    float64
	cost     Currency
	quantity int64
	producer string
	weight   float64
}

func ProductDefaultConstructor() *Product {
	p := new(Product)

	p.name = "New Product Model"
	p.price = 299.99
	p.cost = Currency{"USD", 37.75}
	p.quantity = 100
	p.producer = "New Manufacturer"
	p.weight = 500.5

	return p
}

func ProductConstructor(name string, price float64, cost Currency,
	quantity int64, producer string, weight float64) *Product {
	p := new(Product)

	p.name = name
	p.price = price
	p.cost = cost
	p.quantity = quantity
	p.producer = producer
	p.weight = weight

	return p
}

func ProductConstructorWithoutCost(name string, price float64,
	quantity int64, producer string, weight float64) *Product {
	p := new(Product)

	p.name = name
	p.price = price
	p.cost = Currency{"USD", 37.75}
	p.quantity = quantity
	p.producer = producer
	p.weight = weight

	return p
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}

func (p *Product) SetCost(cost Currency) {
	p.cost = cost
}

func (p *Product) SetWeight(weight float64) {
	p.weight = weight
}

func (p *Product) SetProducer(producer string) {
	p.producer = producer
}

func (p *Product) SetQuantity(quantity int64) {
	p.quantity = quantity
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) GetCost() Currency {
	return p.cost
}

func (p *Product) GetQuantity() int64 {
	return p.quantity
}

func (p *Product) GetWeight() float64 {
	return p.weight
}

func (p *Product) GetProducer() string {
	return p.producer
}

func (p *Product) GetInfo() string {
	priceStr := fmt.Sprintf("%.2f", p.GetPrice())
	costInfo := p.cost.GetInfo()
	quantityStr := fmt.Sprintf("%d", p.GetQuantity())
	weightStr := fmt.Sprintf("%.2f", p.GetWeight())

	info := "Name: " + p.GetName() + "\n"
	info += "Price: " + priceStr + "\n"
	info += "Cost: " + costInfo + "\n"
	info += "Quantity: " + quantityStr + "\n"
	info += "Producer: " + p.GetProducer() + "\n"
	info += "Weight: " + weightStr + "\n"

	return info
}

func (p *Product) GetPriceIn() float64 {
	cost := p.GetCost()
	exRate := cost.GetExRate()

	return exRate * p.GetPrice()
}

func (p *Product) GetTotalPrice(isUah bool) float64 {
	if isUah {
		return p.GetPriceIn() * float64(p.GetQuantity())
	}

	return p.GetPrice() * float64(p.GetQuantity())
}

func (p *Product) GetTotalWeight() float64 {
	return p.GetWeight() * float64(p.GetQuantity())
}
