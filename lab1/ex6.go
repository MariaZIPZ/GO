package main

import "fmt"

func main() {
	var x, y, z uint8

	x = 9
	y = 28
	z = x

	fmt.Println("Битовые операции")

	fmt.Printf("1) ^x      = ^(%d)      = ^(%.8b)            = %.8b = %d\n", x, x, ^x, ^x)
	fmt.Printf("2) x << 2  = (%d << 2)  = (%.8b << 2)        = %.8b = %d\n", x, x, x<<2, x<<2)
	fmt.Printf("3) x >> 2  = (%d >> 2)  = (%.8b >> 2)        = %.8b = %d\n", x, x, x>>2, x>>2)
	fmt.Printf("4) x & y   = (%d & %d)  = (%.8b & %.8b)  = %.8b = %d\n", x, y, x, y, x&y, x&y)
	fmt.Printf("5) x | y   = (%d | %d)  = (%.8b | %.8b)  = %.8b = %d\n", x, y, x, y, x|y, x|y)
	fmt.Printf("6) x ^ y   = (%d ^ %d)  = (%.8b ^ %.8b)  = %.8b = %d\n", x, y, x, y, x^y, x^y)
	fmt.Printf("7) x &^ y  = (%d &^ %d) = (%.8b &^ %.8b) = %.8b = %d\n", x, y, x, y, x&^y, x&^y)
	fmt.Printf("8) x %% y   = (%d %% %d)  = (%.8b %% %.8b)  = %.8b = %d\n", x, y, x, y, x%y, x%y)

	fmt.Println("\nБитовые операции с присваиванием")

	x = z
	x &= y
	fmt.Printf("9) x &= y   = (%d &= %d)  = (%.8b &= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x |= y
	fmt.Printf("10) x |= y   = (%d |= %d)  = (%.8b |= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x ^= y
	fmt.Printf("11) x ^= y   = (%d ^= %d)  = (%.8b ^= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x &^= y
	fmt.Printf("12) x &^= y  = (%d &^= %d) = (%.8b &^= %.8b) = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x %= y
	fmt.Printf("13) x %%= y   = (%d %%= %d)  = (%.8b %%= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)

	// Завдання 1. Пояснити результати операцій
	fmt.Println("\nЗавдання 1. Пояснити результати операцій")

	var explanation string = "\t1. Побітове \"НІ\": Змінюємо кожен біт - 1 на 0, а 0 на 1.\n" +
		"\t2. Зсув вліво на два розряди: Переміщуємо кожен біт вліво на дві позиції, заповнюючи нулями справа.\n" +
		"\t3. Зсув вправо на два розряди: Переміщуємо кожен біт вправо на дві позиції, заповнюючи нулями зліва.\n" +
		"\t4. Побітове \"І\": Результат - 1, якщо обидва відповідні біти рівні 1, в іншому випадку - 0.\n" +
		"\t5. Побітове \"АБО\": Результат - 1, якщо хоча б один з відповідних бітів рівний 1.\n" +
		"\t6. Побітове \"виключне АБО\": Результат - 1, якщо кількість одиниць відповідних бітів непарна.\n" +
		"\t7. Побітове \"І з НІ\": Кожен біт y змінюється протилежним, потім обчислюється \"І\" з відповідним бітом x.\n" +
		"\t8. Отримання залишку від ділення: Визначення остачі від ділення двох чисел (наприклад, залишок від ділення 9 на 28).\n" +
		"\t9 - 13: Подібні операції до вищенаведених, але результат присвоюється першому операнду."

	fmt.Println("Пояснення:\n", explanation)
}