// Задача 2. Проверка уровня заряда аккумулятора на морозе

// При низких температурах аккумулятор автомобиля теряет ёмкость. Нужно проверить, сможет ли он завести машину.

// Входные данные:

// Температура воздуха (в °C)
// Ток холодной прокрутки аккумулятора (CCA) — значение в амперах
// Минимальный ток стартера для запуска — в амперах

// Логические переменные:

// isTooCold — если температура ≤ -25
// isNotEnoughPower — если actualCCA < минимального тока
// Вывод:

// Всегда выводите рассчитанный actualCCA.
// Если isTooCold — выведите «Внимание: очень холодно, риск замерзания аккумулятора!»
// Если isNotEnoughPower — выведите «Недостаточно тока для запуска двигателя.»

package main

import "fmt"

func main() {
	fmt.Println("Введите темперауру воздуха в °C:")

	var airTemperature int 
	fmt.Scanln(&airTemperature)

	const minAirTemperature = -89
	const maxAirTemperature = 71

	if airTemperature < minAirTemperature || airTemperature > maxAirTemperature {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Введите ток холодной прокрутки в амперах:")

	var coldCrankingCurrent int 
	fmt.Scanln(&coldCrankingCurrent)

	fmt.Println("Введите минимальный ток стартера для запуска в амперах:")

	var minStarterCurrent int
	fmt.Scanln(&minStarterCurrent)


}