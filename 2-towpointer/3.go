package towpointer

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	len := len(nums)
	result := make([][]int, 0)

	if len < 3 || nums == nil {
		return result
	}
	for i := 0; i < len; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}
	}
	return result
}

//定义三个指针
