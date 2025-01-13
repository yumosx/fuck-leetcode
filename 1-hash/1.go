package hash

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, i := range nums {
		if value, ok := m[target-i]; ok {
			return []int{value, index}
		} else {
			m[i] = index
		}
	}
	return []int{-1, -1}
}
