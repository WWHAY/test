package stack

func removeDuplicates(S string) string {
	res := []byte{}
	for i := range S {
		if len(res) != 0 && res[len(res)-1] == S[i] {
			res = res[:len(res)-1]
		} else {
			res = append(res, S[i])
		}
	}
	return string(res)
}

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	num := 0
	preSign := '+'
	var stack []int
	for v := range s {
		isDigit := '0' <= v && v <= '9'
		if isDigit {
			num = num*10 + int(v)
		}
		if !isDigit && v != ' ' {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0

		}

	}

	res := 0
	for v := range stack {
		res += v
	}
	return res
}
