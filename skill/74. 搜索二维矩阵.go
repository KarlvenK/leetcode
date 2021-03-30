package main

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	x, y := n-1, m-1
	for {
		if target == matrix[x][y] {
			return true
		}
		if target > matrix[x][y] {
			return false
		}
		if x >= 1 && matrix[x-1][y] >= target {
			x -= 1
		} else {
			if y >= 1 {
				y -= 1
			} else {
				x -= 1
				y = m - 1
				if x < 0 {
					return false
				}
			}
		}
	}
	return false
}

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

solution:
	从最后一个座标开始找，如果上一行同一列的元素大于等于target则直接行数-1
	否则targe要么不存在要么就在这一行，由于我们是从大往小找的，如果出现matrix元素小于target，那么肯定不存在target
*/
