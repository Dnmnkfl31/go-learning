// Задача 4: Разница между len и cap
// Создать пустой слайс с помощью make([]int, 3, 5).
// Вывести слайс, его длину и емкость. Добавить в него 3 элемента через append.
// Вывести слайс, его длину и емкость.
// Что изменилось и почему?

package main

import "fmt"

func main() {
	s := make([]int, 3, 5)

	fmt.Println("cap =", cap(s), "len =", len(s))

	s = append(s, 1, 2, 3)

	fmt.Println("slice:", s)
	fmt.Println("cap =", cap(s))
	fmt.Println("len =", len(s))
}