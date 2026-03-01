// Задача 1. Проверка готовности лыжника к забегу

// Необходимо написать программу, которая помогает лыжнику подготовиться к тренировке на свежем воздухе. 
// Программа должна оценивать погодные условия и длину трассы и выдавать рекомендации: можно ли сегодня безопасно выходить на пробежку и какую экипировку выбрать.

// Входные данные:

// Температура воздуха (°C)
// Скорость ветра (м/с)
// Длина трассы (в километрах)
// Условия:

// Если температура ниже -20°C, слишком холодно для тренировки.
// Если скорость ветра больше 15 м/с, слишком сильный ветер — нельзя выходить.
// Если длина трассы больше 10 км и температура ниже -10°C, рекомендуется надеть дополнительную одежду.
// Вывод:

// Всегда выводите все три показателя: температуру, скорость ветра, длину трассы.
// Если слишком холодно — сообщите, что слишком холодно для тренировки.
// Если ветер слишком сильный — сообщите, что нельзя выходить из-за ветра.
// Если трасса длинная и холодно — напомните про дополнительную одежду.
// Если всё нормально — напишите, что можно начинать тренировку.

package main

import "fmt"

func main() {
	fmt.Println("Введите температуру воздуха:")

	var airTemperature int
	fmt.Scanln(&airTemperature)

	const minAirTemperature = -71
	const maxAirTemperature = 80

	if airTemperature < minAirTemperature || airTemperature > maxAirTemperature {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Введите скорость ветра в м/с:")

	var windSpeed int
	fmt.Scanln(&windSpeed)

	const minWindSpees = 0
	const maxWindSpeed = 50

	if windSpeed < minWindSpees || windSpeed > maxWindSpeed {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Введите длину трассы в км:")

	var lengthTrack int
	fmt.Scanln(&lengthTrack)

	const minLengthTrack = 0

	if lengthTrack < minLengthTrack {
		fmt.Println("Ошибка")
		return
	}

	isTooCold := airTemperature < -20
	isWindStrong := windSpeed > 15
	isRouteLong := lengthTrack > 10 && airTemperature < -10

	fmt.Println("Температура воздуха:", airTemperature)
	fmt.Println("Скорость ветра:", windSpeed)
	fmt.Println("Длина трассы:", lengthTrack)

	if isTooCold {
		fmt.Println("Слищком холодно Оставайтесь дома")

	} else if isWindStrong {
		fmt.Println("Нельзя выходить на улицу из-за ветра")

	} else if isRouteLong {
		fmt.Println("Не забудьте надеть доп одежду")

	} else {
		fmt.Println("Все в порядке. Удачной тренировки")
	}
}