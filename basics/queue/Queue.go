package queue

// Queue A FIFO queue.
type Queue []int

// Push Pushes the element into the queue.
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop Pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
