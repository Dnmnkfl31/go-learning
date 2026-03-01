// Задача 2. Разрешение поездки на самокате

// Условие: Пользователь вводит:

// Возраст
// Есть ли шлем (yes / no)
// Идёт ли дождь (yes / no)
// Создайте логические переменные:

// isAdult   // true, если возраст 18 или старше
// hasHelmet // true, если шлем есть
// isRaining // true, если идет дождь
// Создайте итоговую переменную для if:

// canRideScooter // true, если пользователь взрослый, в шлеме и погода не дождливая
// Вывод: Можете ехать или Поездка запрещена

package main

import "fmt"

func main() {
	fmt.Println("Введите ваш возраст:")

	var age int
	fmt.Scanln(&age)

	fmt.Println("У вас есть шлем? (y / n)")

	var helmet string
	fmt.Scanln(&helmet)

	fmt.Println("На улице идет дождь? (y / n)")

	var rain string
	fmt.Scanln(&rain)

	isAdult := age >= 18
	hasHelmet := helmet == "yes"
	isRaining := rain == "yes"

	canRideScooter := isAdult && hasHelmet && !isRaining

	if canRideScooter {
		fmt.Println("Можете ехать")

	} else {
		fmt.Println("Поездка запрещена")
	}
}