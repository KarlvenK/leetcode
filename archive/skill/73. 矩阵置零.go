package main

func main() {

}

func setZeroes(matrix [][]int) {
	row := len(matrix)
	column := len(matrix[0])
	markRow, markColumn := make([]byte, row), make([]byte, column)
	for i := 0; i < row; i++ {
		markRow[i] = '1'
	}
	for i := 0; i < column; i++ {
		markColumn[i] = '1'
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == 0 && (markRow[i] == '1' || markColumn[j] == '1') {
				markColumn[j] = '0'
				markRow[i] = '0'
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if markRow[i] == '0' || markColumn[j] == '0' {
				matrix[i][j] = 0
			}
		}
	}
}
