package towpointer

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	//不断的交换两个指针
	for right < n {
		if nums[right] != 0 {
			//不断交换两个值
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}
}
