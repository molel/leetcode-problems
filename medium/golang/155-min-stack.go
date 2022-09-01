package golang

type MinStack struct {
	currentNode *StackNode
}
type StackNode struct {
	value    int
	minValue int
	previous *StackNode
}

func Constructor() MinStack {
	return MinStack{&StackNode{}}
}

func (ms *MinStack) Push(val int) {
	if ms.currentNode.previous != nil {
		newNode := &StackNode{value: val, minValue: min(val, ms.currentNode.previous.value), previous: ms.currentNode}
		ms.currentNode = newNode
	} else {
		ms.currentNode.value = val
		ms.currentNode.minValue = val
	}
}

func (ms *MinStack) Pop() {
	ms.currentNode = ms.currentNode.previous
}

func (ms *MinStack) Top() int {
	return ms.currentNode.value
}

func (ms *MinStack) GetMin() int {
	return ms.currentNode.minValue
}
