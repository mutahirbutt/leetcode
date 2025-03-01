package leetcode

func isValid(s string) bool {

	match := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false
		} else {
			if v, ok := match[s[i]]; ok {
				if v != stack[len(stack)-1] {
					return false
				}
			}
			//pop
			stack = stack[:len(stack)-1]

		}

	}

	return len(stack) == 0

}
