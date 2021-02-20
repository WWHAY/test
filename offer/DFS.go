package offer

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if search(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	if board[i][j] == word[k] {
		temp := board[i][j]
		board[i][j] = ' '
		if search(board, i, j+1, k+1, word) ||
			search(board, i, j-1, k+1, word) ||
			search(board, i+1, j, k+1, word) ||
			search(board, i-1, j, k+1, word) {
			return true
		} else {
			board[i][j] = temp
		}
	}
	return false
}

func movingCount(m int, n int, k int) int {
	board := make([][]int, m+1)
	for i := range board {
		board[i] = make([]int, n+1)
	}
	return find(board, m, n, 0, 0, k)
}

func find(board [][]int, m, n, i, j, k int) int {
	if i < 0 || j < 0 || i > m || j > n || (add(i)+add(j) > k) {
		return 0
	}
	sum := 1
	sum += find(board, m, n, i+1, j, k)
	sum += find(board, m, n, i-1, j, k)
	sum += find(board, m, n, i, j+1, k)
	sum += find(board, m, n, i, j-1, k)
	return sum
}

func add(n int) int {
	if n > 0 {
		return n/10 + n%10
	}
	return 0
}
