package main

import (
	"fmt"
)

// Обход массива по спирали от центра влево - вниз - вправо - вверх
func spiralPrint(size int, matrix [][]int) {
	j := size / 2 //устанавливаем итераторы на центральный элемент массива

	i := size / 2
	n := 0

	for true {
		n++
		for k := 0; k < n; k++ { ////выводим n элементов со сдвигом влево
			fmt.Print(matrix[i][j])
			j--
		}
		//точка выхода из цикла
		if n == size { //если сдвиг равен размерности массива
			break //выход
		}
		for k := 0; k < n; k++ { ////выводим n элементов со сдвигом вниз
			fmt.Print(matrix[i][j])
			i++
		}
		n++                      //опять увеличиваем сдвиг
		for k := 0; k < n; k++ { //выводим n элементов со сдвигом вправо
			fmt.Print(matrix[i][j])
			j++
		}

		for k := 0; k < n; k++ { //выводим n элементов со сдвигом вверх
			fmt.Print(matrix[i][j])
			i--
		}
	}
}

func main() {
	var n int
	fmt.Printf("Введите число N для матрицы 2n-1 x 2n-1: ")
	fmt.Scanf("%d", &n)

	sizeMatrix := 2*n - 1
	// выделем под массив простарнство
	matrix := make([][]int, sizeMatrix)
	for i := range matrix {
		matrix[i] = make([]int, sizeMatrix)
	}

	fmt.Printf("Введите матрицу %d x %d", sizeMatrix, sizeMatrix)

	for i := 0; i < sizeMatrix; i++ {
		for j := 0; j < sizeMatrix; j++ {
			_, err := fmt.Scanf("%d", &matrix[i][j])
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Println("Введена матрица вида: ", matrix)

	spiralPrint(sizeMatrix, matrix)
}
