package offer

func replaceSpace(s string) string {
	var strB []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			strB = append(strB, '%', '2', '0')
		} else {
			strB = append(strB, s[i])
		}
	}
	return string(strB)
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	result := 0
	switch b {
	case 0:
		result = remainder(3, a)
	case 1:
		result = remainder(3, a-1) * 4
	default:
		result = remainder(3, a) * 2
	}
	return result % 1000000007
}

func remainder(x, a int) int {
	rem := 1
	for i := 0; i < a; i++ {
		rem = (rem * x) % 1000000007
	}
	return rem
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	var res float64 = 1.0
	if n > 0 {
		for i := 0; i < n; i++ {
			res *= x
		}
	} else {
		m := (-1) * n
		for i := 0; i < m; i++ {
			res *= (1 / x)
		}
	}
	return res
}
