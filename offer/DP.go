package offer

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+2)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]

}

//优化
func numWays1(n int) int {
	a, b := 1, 1
	for i := 1; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = 1
	for j := 1; j < n; j++ {
		dp[j] = grid[0][j]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = max(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
