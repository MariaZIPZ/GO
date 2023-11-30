package main

import (
	"fmt"
)

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведение типов\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання 1. Створіть 2 змінні різних типів. Виконуйте арифметичні операції. Результат вивести
	fmt.Println("\nЗавдання 1. Створіть 2 змінні різних типів. Виконуйте арифметичні операції. Результат вивести")

	var first int16 = -2022
	var second float32 = 2.15
	fmt.Println("first =", first)
	fmt.Println("second =", second)

	unaryMinus := -second
	unaryPlus := +first
	fmt.Println("\n+first", unaryPlus)
	fmt.Println("-second", unaryMinus)

	sum := float32(first) + second
	sub := float32(first) - second
	mul := first * int16(second)
	div := float32(first) / second
	mod := first % int16(second)

	fmt.Println("first + second =", sum)
	fmt.Println("first - second =", sub)
	fmt.Println("first * second =", mul)
	fmt.Println("first / second =", div)
	fmt.Println("first % second =", mod)
}
