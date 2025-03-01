package leetcode

import "strconv"

/*
n Seq
1 1
2 11
3 21
4 1211
5 111221

1 2 1 1
^

*/

func loop(n int) string {
	if n == 1 {
		return "1"
	}

	ans := count("1")
  
	for i := n-2; i > 0; i-- {
		ans = count(ans)
	}
	return ans
}

func count(s string) string {
	bString := []byte(s)
	n := len(bString)
	temp := [][]byte{}
	counter := byte(0)
	prev := bString[0]

	if n == 1 {
		counter++
		return strconv.Itoa(int(counter)) + string(prev)
	}

	for i := range bString {
		if i == 0 {
			counter++
			continue
		}

		if i < n-1 {
			if prev != bString[i] {
				temp = append(temp, []byte{counter, prev})
				prev = bString[i]
				counter = 1
				continue
			}

			counter++
		}

		if i == n-1 {
			if prev != bString[i] {
				temp = append(temp, []byte{counter, prev})
				temp = append(temp, []byte{1, bString[i]})
				continue
			}
			counter++
			temp = append(temp, []byte{counter, prev})
		}

	}

	var ans string

	for _, v := range temp {
		ans += strconv.Itoa(int(v[0]))
		ans += string(v[1])
	}

	return string(ans)

}
