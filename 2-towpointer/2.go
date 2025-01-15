package towpointer

func maxArea(height []int) int {
	// 定义两个指针
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		//计算对应的面积
		area = max(area, (right-left)*min(height[right], height[left]))
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return area
}
