package _4_SpiralMatrix

// 按层遍历
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		index                    = 0
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	for left <= right && top <= bottom {
		// top: [top][left] => [top][right]
		for i := left; i <= right; i++ {
			order[index] = matrix[top][i]
			index++
		}
		// right: [top + 1][right] => [bottom][right]
		for j := top + 1; j <= bottom; j++ {
			order[index] = matrix[j][right]
			index++
		}
		if left < right && top < bottom {
			// bottom: [bottom][right - 1] => [bottom][left]
			for i := right - 1; i > left; i-- {
				order[index] = matrix[bottom][i]
				index++
			}
			// left: [bottom][left] => [top-1][left]
			for j := bottom; j > top; j-- {
				order[index] = matrix[j][left]
				index++
			}
		}
		left++
		right--
		bottom--
		top++
	}
	return order
}
