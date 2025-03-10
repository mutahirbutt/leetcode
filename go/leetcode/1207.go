package leetcode

func uniqueOccurrences(arr []int) bool {
	lookup := map[int]int{}
	for _, v := range arr {
		lookup[v]++
	}

	check := map[int]int{}

	for _, v := range lookup {
		check[v]++
	}

	return len(lookup) == len(check)
}
