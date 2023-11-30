package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Синонимы целых типов\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")

	//Завдання. 1. Визначити розрядність ОС
	fmt.Println("\nЗавдання. 1. Визначити розрядність ОС")

	var bitRate byte

	if math.MaxInt == math.MaxInt64 {
		bitRate = 64
	} else if math.MaxInt == math.MaxInt32 {
		bitRate = 32
	}

	fmt.Println("Розрядність системи:", bitRate)
}
