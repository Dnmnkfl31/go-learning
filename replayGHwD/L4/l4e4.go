// Задача 4. Заказ в магазине
// Условие: Пользователь вводит количество яблок, бананов и апельсинов, 
// которые хочет купить. Если общее количество фруктов больше 20 — выводим: "Слишком большой заказ" 
// Если все три вида фруктов заказаны (т.е. ни один не равен нулю), — "Заказ принят" Иначе — "Добавьте хотя бы один из каждого фрукта"

package main

import "fmt"

func main() {
	fmt.Println("Введите кол-во яблок:")

	var appleCount int
	fmt.Scanln(&appleCount)

	fmt.Println("Введите кол-во бананов:")

	var bananaCount int
	fmt.Scanln(&bananaCount)

	fmt.Println("Введите кол-во апельсинов:")

	var orangeCount int
	fmt.Scanln(&orangeCount)

	sum := appleCount + bananaCount + orangeCount

	if sum > 20 {
		fmt.Println("Слишком большой заказ")

	} else if appleCount != 0 && bananaCount != 0 && orangeCount != 0 {
		fmt.Println("Заказ принят")

	} else {
		fmt.Println("Добавьте хотя бы один фрукт")
	}
}