// Задача 3. Авторизация
// Условие: Пользователь вводит имя пользователя и пароль. Если имя admin и пароль 1234, 
// то выводим "Вход выполнен" Если имя admin, но пароль неверный — "Неверный пароль" 
// Во всех остальных случаях — "Неизвестный пользователь"

package main

import "fmt"

func main() {
	fmt.Println("Введите имя пользователя:")

	var name string
	fmt.Scanln(&name)

	fmt.Println("Введите пароль:")

	var password int
	fmt.Scanln(&password)

	if name == "admin" && password == 1234 {
		fmt.Println("Вход выполнен")

	} else if name == "admin" && password != 1234 {
		fmt.Println("Неверный пароль")

	} else {
		fmt.Println("Неизвестный пользователь")
	}
}