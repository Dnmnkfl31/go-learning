package main

import (
	"fmt"
	"math/rand"
)
 
func main() {

	fmt.Println("Добро пожаловать в казино! Стоимость спина: 1000₽")

	fmt.Println("Введите депозит кратный стоимости спина: ")

	var deposit int
	fmt.Scanln(&deposit)

	const minDeposit = 1000

	if deposit % 1000 != 0 {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Ваш депозит =", deposit)

	const (
		diamond = 0
		lemon = 1
		cherry = 2
	)

	const start = 1
	const exit = 2

	for {
		fmt.Println("Введите команду: 1 - играть; 2 - выйти.")

		var gameLoginCommand int
		fmt.Scanln(&gameLoginCommand)

		if gameLoginCommand == exit {
			fmt.Println("Всего доброго!")
			return
		}

		if gameLoginCommand != start && gameLoginCommand != exit {
			fmt.Println("Ошибка")
			continue
		}

		pic_1 := rand.Intn(3)
		pic_2 := rand.Intn(3)
		pic_3 := rand.Intn(3)

		if pic_1 == diamond {
		fmt.Print("[ алмаз | ")
		} else if pic_1 == lemon {
		fmt.Print("[ лимон | ")
		} else if pic_1 == cherry {
		fmt.Print("[ вишенка | ")
		}

		if pic_2 == diamond {
		fmt.Print("алмаз")
		} else if pic_2 == lemon {
		fmt.Print("лимон")
		} else if pic_2 == cherry {
		fmt.Print("вишенка")
		}

		if pic_3 == diamond {
		fmt.Print(" | алмаз ]")
		} else if pic_3 == lemon {
		fmt.Print(" | лимон ]")
		} else if pic_3 == cherry {
		fmt.Print(" | вишенка ]")
		}

		deposit -= 1000

		if pic_1 == pic_2 || pic_1 == pic_3 || pic_2 == pic_3 {
			deposit += 500
		}

		if pic_1 == pic_2 && pic_1 == pic_3 && pic_2 == pic_3 {
			deposit += 3000
		}

		fmt.Println(" Ваш депозит =", deposit)

		if deposit < minDeposit {
			return
		}
	}
}