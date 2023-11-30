package main

import (
	"fmt"
	"reflect"
)

func main() {
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	//Завдання. 1. Вивести типи всіх змінних
	fmt.Println("\nЗавдання. 1. Вивести типи всіх змінних")

	// Використовуючи fmt.Printf та специфікатор %T:
	fmt.Printf("Name = userinit8, Value = %d, Type = %T\n", userinit8, userinit8)
	fmt.Printf("Name = userinit16, Value = %d, Type = %T\n", userinit16, userinit16)

	// Використовуючи reflect та методи reflect.TypeOf() та reflect.ValueOf()
	fmt.Println("Name = userinit64, Value =", reflect.ValueOf(userinit64), "Type =", reflect.TypeOf(userinit64))
	fmt.Println("Name = userautoinit, Value =", reflect.ValueOf(userautoinit), "Type =", reflect.TypeOf(userautoinit))

	//Завдання 2. Присвоїти змінної intVar змінні userinit16 і userautoinit. Результат вивести.
	fmt.Println("\nЗавдання 2. Присвоїти змінної intVar змінні userinit16 і userautoinit. Результат вивести.")

	intVar = int(userinit16) // Виконуємо приведення типу uint16 до int
	fmt.Println("Name = intVar, Value =", reflect.ValueOf(intVar), "Type =", reflect.TypeOf(intVar))

	intVar = userautoinit
	fmt.Println("Name = intVar, Value =", reflect.ValueOf(intVar), "Type =", reflect.TypeOf(intVar))
}