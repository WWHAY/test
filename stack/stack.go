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
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	var ans int
	for _, v := range stack {
		ans += v
	}
	return ans
}
