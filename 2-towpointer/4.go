package towpointer

type Stack struct {
	Len   int
	top   int
	value []int
}

func NewStack() *Stack {
	return &Stack{Len: 0, value: make([]int, 0)}
}

func (st *Stack) Pop() int {
	v := st.value[st.Len]
	st.Len--
	st.Top = st.value[st.Len]
	return v
}

func (st *Stack) Push(value int) {
	st.top = value
	st.Len++
	st.value = append(st.value, value)
}

func (st *Stack) GetTop() int {
	if st.Len == 0 {
		return -1
	}
	return st.top
}

func trap(height []int) int {
	st := NewStack()
	last := 0
	res := 0

	for i := 0; i < len(height); i++ {
		for st.Len != 0 && height[i] > height[st.GetTop()] {
			res += (i - st.GetTop() - 1) * (height[st.GetTop()] - last)
			last = height[st.GetTop()]
			st.Pop()
		}

		if st.Len != 0 {
			res += (i - st.GetTop() - 1) * (height[i] - last)
		}

		st.Push(i)
	}

	return res
}
