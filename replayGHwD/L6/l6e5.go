// Задача 5. Можно ли регистрироваться в сервисе

// Пользователь вводит:

// Возраст
// Принял ли пользователь условия использования (yes / no)
// Подтвердил ли пользователь email (yes / no)
// Создайте логические переменные:

// isOldEnough       // true, если возраст ≥ 13
// hasAcceptedTerms  // true, если введено "yes"
// hasConfirmedEmail // true, если введено "yes"
// Создайте итоговую переменную для if:

// canRegister // true, если выполнены все 3 условия
// Вывод: Регистрация выполнена или Регистрация невозможна

package main

import "fmt"

func main() {
	fmt.Println("Введите ваш возраст:")

	var age int
	fmt.Scanln(&age)

	const minAge = 5
	const maxAge = 100

	if age < minAge || age > maxAge {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Вы приняли пользовательское соглашение? (y / n)")

	var acceptedTerms string
	fmt.Scanln(&acceptedTerms)

	const yes = "yes"
	const no = "no"

	if acceptedTerms != yes && acceptedTerms != no {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Вы подтсвердили email? (y / n)")

	var confirmedEmail string
	fmt.Scanln(&confirmedEmail)

	if confirmedEmail != yes && confirmedEmail != no {
		fmt.Println("Ошибка")
		return
	}

	isOldEnough := age >= 13
	hasAcceptedTerms := acceptedTerms == yes
	hasConfirmedEmail := confirmedEmail == yes

	canRegister := isOldEnough && hasAcceptedTerms && hasConfirmedEmail

	if canRegister {
		fmt.Println("Регистрация выполнена")

	} else {
		fmt.Println("Регистрация невозможна")
	}
}