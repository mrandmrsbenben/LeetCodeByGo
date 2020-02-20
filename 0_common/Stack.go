package common

type Stack struct {
	data []int
}

func StackConstructor() Stack {
	return Stack{make([]int, 0)}
}

func (t *Stack) Push(n int) {
	t.data = append(t.data, n)
}

func (t *Stack) Pop() int {
	if len(t.data) == 0 {
		return -1
	}

	top := t.data[len(t.data)-1]
	t.data = t.data[0 : len(t.data)-1]
	return top
}
