package ui

import (
	"errors"
	"fmt"
	"lab5/packages/business"
	"os"
)

func ReadProductsArray() ([]business.Product, error) {
	var products []business.Product

	var size int
	fmt.Print("Enter the number of elements in the products array: ")
	_, err := fmt.Fscan(os.Stdin, &size)

	if err != nil {
		return nil, errors.New("incorrect array length value")
	}

	for i := 0; i < size; i++ {
		fmt.Println("\nFilling information about product #", i+1)

		item, err := readProduct()
		if err != nil {
			fmt.Println("Error during input! Default constructor selected.")
			item = business.ProductDefaultConstructor()
		}
		products = append(products, *item)
	}

	return products, nil
}

func readProduct() (*business.Product, error) {
	fmt.Println("Choose the type of product structure constructor: \n1) With all fields \n2) Default currency \n3) Default")

	var constructorType int64
	fmt.Print("Enter the selected option: ")
	_, err := fmt.Fscan(os.Stdin, &constructorType)

	var name string
	var price float64
	var cost *business.Currency
	var quantity int64
	var producer string
	var weight float64

	switch constructorType {
	case 1:
		name, price, quantity, producer, weight, err = readProductDetails()
		if err != nil {
			return nil, err
		}

		cost, err = readCurrency()
		if err != nil {
			return nil, err
		}

		return business.ProductConstructor(name, price, *cost, quantity, producer, weight), nil
	case 2:
		name, price, quantity, producer, weight, err = readProductDetails()
		if err != nil {
			return nil, err
		}

		return business.ProductConstructorWithoutCost(name, price, quantity, producer, weight), nil
	case 3:
		return business.ProductDefaultConstructor(), nil
	default:
		fmt.Println("Error! Incorrect option selected. Default constructor chosen.")
		return business.ProductDefaultConstructor(), nil
	}
}

func readCurrency() (*business.Currency, error) {
	fmt.Println("Choose the type of currency structure constructor: \n1) With name and exchange rate \n2) Only exchange rate \n3) Default")

	var constructorType int64
	fmt.Print("Enter the selected option: ")
	_, err := fmt.Fscan(os.Stdin, &constructorType)

	var name string
	var exRate float64

	switch constructorType {
	case 1:
		fmt.Print("Enter the currency name: ")
		_, err = fmt.Fscan(os.Stdin, &name)

		fmt.Print("Enter the exchange rate: ")
		_, err = fmt.Fscan(os.Stdin, &exRate)

		if err != nil {
			fmt.Println("Error during input! Default constructor selected.")
			return business.CurrencyDefaultConstructor(), err
		}

		return business.CurrencyConstructor(name, exRate), nil
	case 2:
		fmt.Print("Enter the exchange rate: ")
		_, err = fmt.Fscan(os.Stdin, &exRate)

		if err != nil {
			fmt.Println("Error during input! Default constructor selected.")
			return business.CurrencyDefaultConstructor(), err
		}

		return business.CurrencyConstructorWithExRate(exRate), nil
	case 3:
		return business.CurrencyDefaultConstructor(), nil
	default:
		fmt.Println("Error! Incorrect option selected. Default constructor chosen.")
		return business.CurrencyDefaultConstructor(), nil
	}
}

func readProductDetails() (string, float64, int64, string, float64, error) {
	var name string
	var price float64
	var quantity int64
	var producer string
	var weight float64
	var err error

	fmt.Print("Enter the product name: ")
	_, err = fmt.Fscan(os.Stdin, &name)

	fmt.Print("Enter the product price: ")
	_, err = fmt.Fscan(os.Stdin, &price)

	fmt.Print("Enter the product quantity: ")
	_, err = fmt.Fscan(os.Stdin, &quantity)

	fmt.Print("Enter the product producer name: ")
	_, err = fmt.Fscan(os.Stdin, &producer)

	fmt.Print("Enter the product weight: ")
	_, err = fmt.Fscan(os.Stdin, &weight)

	if err != nil {
		fmt.Println("Error during input! Default values will be used.")
		return "", 0, 0, "", 0, err
	}

	return name, price, quantity, producer, weight, nil
}
