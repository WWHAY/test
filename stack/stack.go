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
