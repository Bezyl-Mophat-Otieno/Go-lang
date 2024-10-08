package lesson_six

import "fmt"

func Arrays(){
	var numbers [5]int;
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)
	var matrix [3][3]int 
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	fmt.Println(matrix)

		for i:=0; i<3; i++ {
		  for j:=0; j<3; j++ {
            fmt.Printf("%d ", matrix[i][j])
		  }
		          fmt.Println()
		}
}