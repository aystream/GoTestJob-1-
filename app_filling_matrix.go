package main

import (
	"fmt"
	"math"
)

/*
 Заполняет по спирали матрицу:  влево - вниз - вправо - вверх
*/

// влево - вниз - вправо - вверх
func SpiralFillingMatrix(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := int(math.Min(math.Min(float64(i), float64(j)), math.Min(float64(n-1-i), float64(n-1-j))))

			// Для верхней правой половины
			if i <= j {
				fmt.Printf("%d\t ", (n-2*x)*(n-2*x)-(i-x)-(j-x))
			} else { // для нижней левой половины
				fmt.Printf("%d\t ", (n-2*x-2)*(n-2*x-2)+(i-x)+(j-x))
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	n := 4
	size := 2*n - 1
	SpiralFillingMatrix(size)
}
