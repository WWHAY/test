package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		m := len(matrix[i])
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] < target {
				break
			}
		}
	}
	return false
}
