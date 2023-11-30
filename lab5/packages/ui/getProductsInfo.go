package ui

import "lab5/packages/business"

func GetProductsInfo(products []business.Product) (min business.Product, max business.Product) {
	min = products[0]
	max = products[0]

	for i := 1; i < len(products); i++ {
		if products[i].GetPriceIn() < min.GetPriceIn() {
			min = products[i]
		} else if products[i].GetPriceIn() > max.GetPriceIn() {
			max = products[i]
		}
	}

	return min, max
}
