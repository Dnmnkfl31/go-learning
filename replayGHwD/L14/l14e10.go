// Задача 10. Сортировка строк
// Создать слайс строк []string{"go", "java", "python", "csharp", "ruby"}. 
// Отсортировать его в алфавитном порядке с помощью sort.Strings. Вывести результат.

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"go", "java", "python", "csharp", "ruby"}

	fmt.Println(s)

	sort.Strings(s)

	fmt.Println(s)
}