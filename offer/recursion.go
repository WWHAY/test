package offer

func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	sum := n

	sum += sumNums(n - 1)
	return sum
}
