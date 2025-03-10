package leetcode

/*
1 2 3
     ^
*/

func moveZeroes(nums []int) {
	z, r := 0, 0
	n := len(nums)

	for z < n && r < n {
		if nums[z] != 0 {
			z++
			continue
		}

		r = z + 1

		for r < n {
			if nums[r] == 0 {
				r++
				continue
			}

			nums[z], nums[r] = nums[r], nums[z]
			z++
			break
		}
	}
}
