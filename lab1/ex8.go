package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання. 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести
	fmt.Println("Завдання. 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести")
	var defaultComplex64 complex64
	complexNumber := 5683 + 23i
	fmt.Println("defaultComplex:", defaultComplex64, "type: ", reflect.TypeOf(defaultComplex64))
	fmt.Println("complex:", complexNumber, "type: ", reflect.TypeOf(complexNumber))

	runesArray := []rune("🌍😊")
	fmt.Println("runes:", runesArray, "type: ", reflect.TypeOf(runesArray))

	var defaultString string
	word := "Supercalifragilisticexpialidocious"
	fmt.Println("defaultString:", defaultString, "type: ", reflect.TypeOf(defaultString))
	fmt.Println("string:", word, "type: ", reflect.TypeOf(word))

	var defaultBoolean bool
	isTrue := true
	fmt.Println("defaultBoolean:", defaultBoolean, "type: ", reflect.TypeOf(defaultBoolean))
	fmt.Println("boolean:", isTrue, "type: ", reflect.TypeOf(isTrue), "size: ", unsafe.Sizeof(isTrue))

	var defaultUint8 uint8
	numberInt := 127
	fmt.Println("defaultUint8:", defaultUint8, "type: ", reflect.TypeOf(defaultUint8))
	fmt.Println("int:", numberInt, "type: ", reflect.TypeOf(numberInt), "size: ", unsafe.Sizeof(numberInt))

	var defaultFloat32 float32
	fmt.Println("float32:", defaultFloat32, "type: ", reflect.TypeOf(defaultFloat32), "size: ", unsafe.Sizeof(defaultFloat32))

	first, second := 5.2, -60234545.23
	fmt.Println("float64:", first, "type: ", reflect.TypeOf(first), "size: ", unsafe.Sizeof(first))
	fmt.Println("float64:", second, "type: ", reflect.TypeOf(second), "size: ", unsafe.Sizeof(second))
}
