package main

import (
	"fmt"
	functions "lab5/packages/ui"
)

func main() {
	fmt.Println("\nReading products:")
	products, err := functions.ReadProductsArray()

	if err != nil {
		fmt.Println("An error occurred while reading the array.")
	}

	fmt.Println("\nDisplaying information about products:")
	functions.PrintProducts(products)

	min, max := functions.GetProductsInfo(products)
	fmt.Println("Cheapest product:", min.GetName(), fmt.Sprintf("%.2f", min.GetPriceIn()), "UAH")
	fmt.Println("Most expensive product:", max.GetName(), fmt.Sprintf("%.2f", max.GetPriceIn()), "UAH")
}
