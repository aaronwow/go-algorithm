package algorithm

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1

	x, y, num := 0, 0, 1

	for num <= n*n {
		// top
		for y = left; y <= right; y++ {
			matrix[top][y] = num
			num++
		}
		top++

		// right
		for x = top; x <= bottom; x++ {
			matrix[x][right] = num
			num++
		}
		right--

		// bottom
		for y = right; y >= left; y-- {
			matrix[bottom][y] = num
			num++
		}
		bottom--

		// left
		for x = bottom; x >= top; x-- {
			matrix[x][left] = num
			num++
		}
		left++
	}

	return matrix
}
