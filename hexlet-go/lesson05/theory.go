package lesson05

import (
	"errors"
	"fmt"
)

func Multiply(x int, y int) int {
	return x * y
}

// приватная функция, извне не вызвать
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide on zero")
	}

	return x / y, nil
}

func myPrint(msg string) {
	// пакет.функция
	fmt.Println(msg)
}

//	Функция объявляется следующим образом:
//
//	func имя_функции (список_параметров) (типы_возвращаемых_значений){
//	    выполняемые_операторы
//	}