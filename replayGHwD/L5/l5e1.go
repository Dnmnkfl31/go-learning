// Доделать!!!!!

// Задача 1. Проверка допуска к экзамену

// Условие: Пользователь вводит:

// Количество пропущенных занятий
// Средний балл
// Есть ли медицинская справка (yes / no)
// Создайте логические переменные:

// hasTooManyAbsences     // true, если пропущено больше 5 занятий
// hasLowGrade            // true, если средний балл ниже 3.0
// hasMedicalCertificate  // true, если введено "yes"
// Создайте итоговую переменную для if:

// isAllowedToExam // true, если нет слишком большого количества пропусков И балл нормальный ИЛИ есть справка
// Логика допуска:

// Если слишком много пропусков и низкий балл — можно только со справкой.
// Если справки нет — не допустить.
// Всё остальное — допустить.
// Вывод: Допщуен или Не допущен

package main

import "fmt"

func main() {
	fmt.Println("Введите кол-во пропущенных занятий:")

	var countMissedClasses int
	fmt.Scanln(&countMissedClasses)

	fmt.Println("Введите свой средний балл:")

	var average float64
	fmt.Scanln(&average)

	fmt.Println("У вас есть медицинская справка?")

	var medicalCertificate string
	fmt.Scanln(&medicalCertificate)

	hasTooManyAbsences := countMissedClasses > 5
	hasLowGrade := average <= 3.0
	hasMedicalCertificate := medicalCertificate == "да"

	isAllowedToExam := !hasTooManyAbsences && !hasLowGrade && hasMedicalCertificate

	if hasTooManyAbsences && hasLowGrade {
		fmt.Println("Допущен только со справкой")

	} else if !hasMedicalCertificate {
		fmt.Println("Не допустить")

	} 
}