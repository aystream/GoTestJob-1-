package main

import (
	"fmt"
	"strconv"
)

/*
 Рекурсивно заполняет по спирали матрицу:  влево - вниз - вправо - вверх
*/

// возвращает элемент в позиции x, y в спирали размера n
func recursionElement(n, x, y int) int {
	squaring := n * n

	if n%2 == 0 {
		if y == n-1 { // верхняя линия: расчитываем x = n^2 - n
			return squaring - n + x
		}
		if x == 0 { // правая колонка: расчитываем y = n^2 - n - (n - 1)
			return squaring - n - (n - 1) + y
		}
		return recursionElement(n-1, x-1, y) // recursion to odd subspiral
	}

	if y == 0 { // нижняя линия - расчитываем x = n^2 - 1
		return squaring - 1 - x
	}
	if x == n-1 { // левая колонка: расчитываем y = n^2 - n
		return squaring - n - y
	}
	return recursionElement(n-1, x, y-1) // рекурсивно по спирали
}

// влево - вниз - вправо - вверх
func Spiral(n int) {
	str := strconv.Itoa(n*n - 1)
	length := len(str)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			fmt.Printf("%[2]*[1]v ", recursionElement(n, x, y), length)
		}
		fmt.Println()
	}
}

func main() {
	n := 4
	size := 2*n - 1
	Spiral(size)
}
