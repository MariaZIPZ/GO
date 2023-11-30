package ui

import (
	"fmt"
	"lab5/packages/business"
)

func PrintProducts(products []business.Product) {
	for index, item := range products {
		fmt.Println("\nProduct #", index+1)
		PrintProduct(item)
	}
}

func PrintProduct(p business.Product) {
	fmt.Println(p.GetInfo())
}
