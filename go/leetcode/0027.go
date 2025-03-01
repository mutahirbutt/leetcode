package leetcode

/*
3 2 2 3
^
      ^
*/

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val && nums[right] != val {
			tmp := nums[right]
			nums[left] = tmp
			nums[right] = val
			left++
			right--
			continue
		}

		if nums[left] != val {
			left++
			continue
		}

		if nums[left] == val {
			right--
			continue
		}

	}
	return left

}
