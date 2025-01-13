package hash

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	longestStreak := 0

	for num := range m {
		if !m[num-1] { // 检查是否是起始部分
			currentNum := num
			currentStreak := 1

			for m[currentNum+1] {
				currentNum += 1
				currentStreak += 1
			}

			longestStreak = max(longestStreak, currentStreak)
		}
	}

	return longestStreak
}
