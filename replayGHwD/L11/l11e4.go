// 4. Пропускаем пятерки (3)

// Пропускаем пятерки Выведите все числа от 1 до 15, кроме тех, 
// которые делятся на 5. Пример вывода: 1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 13, 14.

package main

import "fmt"

func main() {
	// for i := 1; i <= 15; i++ {
	// 	if i % 5 == 0 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	i := 1

// 	for i < 16 {
// 		if i % 5 == 0{
// 			i++
// 			continue
// 		}
// 		fmt.Println(i)
	
// 	i++
// 	}

for {
	if i % 5 == 0 {
		i++
		continue
	}
	fmt.Println(i)

	i++
	if i >= 15 {
		break
	}
}
}