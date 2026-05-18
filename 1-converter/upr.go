package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

const IMTPower = 2

func main() {
	// Приветственный вывод
	fmt.Println("__Колькулятор массы тела__")

	// Главный цикл программы: спрашиваем, считаем, показываем, повторяем
	for {
		userKg, userHeight := getUserImput()         // запрос данных от пользователя
		IMT, err := calculeteIMT(userKg, userHeight) // расчет ИМТ и проверка
		if err != nil {
			fmt.Println(err)
			continue // если некорректные данные — повторяем цикл
		}

		outputResult(IMT) // вывод результата и категории

		var y string
		fmt.Println("Хотите повторить?(N/Y)")
		fmt.Scan(&y) // ответ пользователя о продолжении
		if y == "N" || y == "n" {
			break // выход из цикла, завершение программы
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculeteIMT(userKg, userHeight float64) (float64, error) {
	// Проверка на корректные значения (рост и вес должны быть больше нуля)
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("Некорректно указано значение, попробуй еще раз! ")
	}

	// Формула ИМТ: вес(кг) / (рост(м)^2)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserImput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	_, e := fmt.Scan(&userHeight)
	if e != nil {
		fmt.Println("Ошибка при вводе роста")
		flushStdin()
	}
	fmt.Print("Введите свой вес: ")
	_, e = fmt.Scan(&userKg)
	if e != nil {
		fmt.Println("Ошибка при вводе веса")
		flushStdin()
	}
	return userKg, userHeight
}

func flushStdin() {
	// При ошибках ввода с fmt.Scan() в stdin может оставаться ненужная часть строки.
	// Эта утилита "съедает" остаток ввода до следующего переноса строки.
	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')
}
