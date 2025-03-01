package leetcode

func tribonacci(n int) int {
	cache := map[int]int{}
	cache[0] = 0
	cache[1] = 1
	cache[2] = 1

	return loop(n, cache)

}

func loop(n int, cache map[int]int) int {

	if n == 0 || n == 1 || n == 2 {
		return cache[n]
	}

	if v, ok := cache[n]; ok {
		return v
	}

	res := loop(n-1, cache) + loop(n-2, cache) + loop(n-3, cache)
	cache[n] = res
	return res

}
