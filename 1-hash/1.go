package hash

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, i := range nums {
		//每次在遍历的时候, 判断一下和当前配置的key
		if value, ok := m[target-i]; ok {
			return []int{value, index}
		} else {
			m[i] = index
		}
	}
	return []int{-1, -1}
}
