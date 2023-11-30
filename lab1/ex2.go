package main

import "fmt"

func main() {
	var defaultInt8 int8
	var defaultInt16 int16
	var defaultInt32 int32
	var defaultInt64 int64
	var defaultInt int

	fmt.Println("Default values (signed): ", defaultInt8, defaultInt16, defaultInt32, defaultInt64, defaultInt)

	var defaultuInt8 uint8
	var defaultuInt16 uint16
	var defaultuInt32 uint32
	var defaultuInt64 uint64
	var defaultuInt uint

	fmt.Println("Default values (unsigned): ", defaultuInt8, defaultuInt16, defaultuInt32, defaultuInt64, defaultuInt)

	//Завдання 1. Створити цілочисельну змінну (результат не відображати)

	//var (
	//first  uint8  = 255       // Максимальне значення для uint8 (максимум: 255)
	//second int16  = 32459     // Знаходиться в межах допустимого для int16
	//third  uint32 = 4294967295 // Максимальне значення для uint32 (максимум: 4294967295)
	//fourth int64  = 9223372036854775807 // Максимальне значення для int64 (максимум: 9223372036854775807)
	//)

	// Закоментовуємо, так як програма, яка містить оголошені зміннні, але ніде не використані не скомпілюється
}
