// Задача 1. Розыгрыш баллов в приложении

// При заказе в приложении доставки еды в выходной день и случайном "везении" (шанс 30%) пользователь получает бонусные баллы.

// Входные данные:

// День недели (Mon, Tue, Wed, Thu, Fri, Sat, Sun)
// Создайте логические переменные:

// isWeekend
// isLuckyToday — 30% шанс (сгенерировать случайное число от 1 до 100, и если оно < 30, то считается, что пользователю повезло)
// Создайте итоговую переменную:

// getsBonus - true, если день выходной и при этом пользователю повезло.
// Если getsBonus = true, бонус случайное число от 1 до 500, иначе 0.

// Выведите: "Вы получили бонус: N" или "В следующий раз повезет!".

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Введите день недели:")

	var dayOfWeek string
	fmt.Scanln(&dayOfWeek)

	const monday = "monday"
	const tuesday = "tuesday"
	const wednesday = "wednesday"
	const thurday = "thurday"
	const friday = "friday"
	const saturday = "saturday"
	const sunday = "sunday"

	if dayOfWeek != monday &&
	   dayOfWeek != tuesday &&
	   dayOfWeek != wednesday &&
	   dayOfWeek != thurday &&
	   dayOfWeek != friday &&
	   dayOfWeek != saturday &&
	   dayOfWeek != sunday {
		
		fmt.Println("Некорректный день недели")
		return
	}

	randomNumber := rand.Intn(100) + 1
	fmt.Println(randomNumber)

	isWeekend := dayOfWeek == friday || dayOfWeek == saturday
	isLuckyToday := 1 <= randomNumber && randomNumber <= 30

	getsBonus := isWeekend && isLuckyToday

	if getsBonus {
		discountAmount := rand.Intn(500) + 1
		fmt.Println("Вы получили бонус: ", discountAmount)

	} else {
		fmt.Println("В след раз повезет!")
	}
}