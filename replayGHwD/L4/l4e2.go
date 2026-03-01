// Задача 2. Проверка скидки
// Условие: Пользователь вводит свой возраст и статус (например: student, retired, none). 
// Если пользователь младше 18 или старше 65, или если он студент — выводим "Скидка применена", иначе — "Скидка не положена"

package main

import "fmt"

func main() {
	fmt.Println("Введите свой возраст:")

	var age int
	fmt.Scanln(&age)

	fmt.Println("Введите свой статус:")

	var status string
	fmt.Scanln(&status)

	if age < 18 || age > 65 {
		fmt.Println("Скидка применена")

	} else if status == "student" {
		fmt.Println("Скидка применена")

	} else {
		fmt.Println("Скидка не полагается")
	}
}