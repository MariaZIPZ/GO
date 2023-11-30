package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("1) first  = ", first)       // false
	fmt.Println("2) second = ", second)      // false
	fmt.Println("3) third  = ", third)       // true
	fmt.Println("4) fourth = ", fourth)      // false
	fmt.Println("5) fifth  = ", fifth, "\n") // true

	fmt.Println("6) !true  = ", !true)        // false
	fmt.Println("7) !false = ", !false, "\n") // true

	fmt.Println("8) true && true   = ", true && true)          // true
	fmt.Println("9) true && false  = ", true && false)         // false
	fmt.Println("10) false && false = ", false && false, "\n") // false

	fmt.Println("11) true || true   = ", true || true)         // true
	fmt.Println("12) true || false  = ", true || false)        // true
	fmt.Println("13) false || false = ", false || false, "\n") // false

	fmt.Println("14) 2 < 3  = ", 2 < 3)        // true
	fmt.Println("15) 2 > 3  = ", 2 > 3)        // false
	fmt.Println("16) 3 < 3  = ", 3 < 3)        // false
	fmt.Println("17) 3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("18) 3 > 3  = ", 3 > 3)        // false
	fmt.Println("19) 3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("20) 2 == 3 = ", 2 == 3)       // false
	fmt.Println("21) 3 == 3 = ", 3 == 3)       // true
	fmt.Println("22) 2 != 3 = ", 2 != 3)       // true
	fmt.Println("23) 3 != 3 = ", 3 != 3, "\n") // false

	//Завдання 1. Пояснити результати операцій
	fmt.Println("Завдання 1. Пояснити результати операцій")

	var explanation string = "Пояснення:\n" +
		"\t 1 та 2: Значення змінних first та second - false, оскільки нульове значення для типу bool = false.\n" +
		"\t 3: Змінну third було ініціалізовано значенням true.\n" +
		"\t 4: Змінна fourth ініціалізована протилежним значенням змінної third (логічне НІ).\n" +
		"\t 5: Змінна fifth було ініціалізовано значенням true.\n" +
		"\t 6 та 7: Логічне НІ змінює значення типу bool на протилежне (true = !false).\n" +
		"\t 8, 9 та 10: При логічному І всі вирази повинні давати значення true, тож true && true = true; true && false = false; false && false = false.\n" +
		"\t 11, 12 та 13: При логічному АБО один з виразів повинен давати значення true, тож true || true = true; true || false = true; false || false = false.\n" +
		"\t 14: Вираз 2 < 3: математично правильно, тому true.\n" +
		"\t 15: Вираз 2 > 3: математично неправильно, тому false.\n" +
		"\t 16: Вираз 3 < 3: математично неправильно, тому false.\n" +
		"\t 17: Вираз 3 <= 3: математично правильно, оскільки 3 ≤ 3, тому true.\n" +
		"\t 18: Вираз 3 > 3: математично неправильно, тому false.\n" +
		"\t 19: Вираз 3 >= 3: математично правильно, оскільки 3 ≥ 3, тому true.\n" +
		"\t 20: Вираз 2 != 3: математично правильно, оскільки 2 ≠ 3, тому true.\n" +
		"\t 21: Вираз 3 != 3: математично неправильно, оскільки 3 = 3, тому false."

	fmt.Println(explanation)
}
