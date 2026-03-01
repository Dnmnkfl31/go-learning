// Задача 4. Разрешение на ночную вечеринку

// Пользователь вводит:

// Возраст
// Разрешение от родителей (yes / no)
// День недели (friday, saturday, sunday, monday, и т.д.)
// Создайте логические переменные:

// isOver18              // true, если возраст ≥ 18
// hasParentalPermission // true, если введено "yes"
// isWeekend             // true, если день — friday или saturday
// Создайте итоговую переменную для if:

// canGoToParty // true, если человек ≥ 18 ИЛИ есть разрешение и это выходной
// Вывод: Можете проходить или Мы не можем вас пропустить

package main

import "fmt"

func main() {
	fmt.Println("Введите ваш возраст:")

	var age int
	fmt.Scanln(&age)

	fmt.Println("У вас есть разрешение от родителей? (y / n)")

	var parentalPermission string
	fmt.Scanln(&parentalPermission)

	fmt.Println("Введите день недели:")

	var numDay string
	fmt.Scanln(&numDay)

	isOver18 := age >= 18
	hasParentalPermission := parentalPermission == "yes"
	isWeekend := numDay == "friday" || numDay == "saturday"

	canGoToParty := isOver18 || hasParentalPermission && isWeekend

	if canGoToParty {
		fmt.Println("Можете проходить")

	} else {
		fmt.Println("Мы не можем вас пропустить")
	}
}