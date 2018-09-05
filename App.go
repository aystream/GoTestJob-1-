package main

import "fmt"

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

	// начиная из центра по спирали: влево - вниз - вправо - вверх
	i := 0
	j := sizeMatrix - 1

	for sizeMatrix != 0 {
		k := 0

		for ok := true; ok; ok = k < sizeMatrix-1 {
			k++
			j--
			fmt.Print(matrix[i][j])
		}

		for k = 0; k < sizeMatrix-1; k++ {
			fmt.Print(matrix[i][j])
			i++
		}
		for k = 0; k < sizeMatrix-1; k++ {
			fmt.Print(matrix[i][j])
			j++
		}
		for k = 0; k < sizeMatrix-1; k++ {
			fmt.Print(matrix[i][j])
			i--
		}
		/*i++
		j++*/
		if sizeMatrix < 2 {
			sizeMatrix = 0
		} else {
			sizeMatrix = sizeMatrix - 2
		}
	}
}
