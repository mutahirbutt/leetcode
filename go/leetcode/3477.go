package leetcode

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	check := map[int]bool{}

	for i := range baskets {
		check[i] = false
	}

	n := len(fruits)
	m := len(check)

	for i := 0; i < n; i++ {
		j := 0
		for j < m {
			if !check[j] {
				if fruits[i] <= baskets[j] {
					check[j] = true
					break
				}
			}
			j++
		}
	}

	counter := 0
	for _, v := range check {
		if !v {
			counter++
		}
	}

	return counter

}
