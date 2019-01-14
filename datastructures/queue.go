package datastructures

// FIFO Queue
type Queue struct {
	queue *LinkedList
}

func (q *Queue) Enqueue(d interface{}) {
	q.queue.Append(d)
}

func (q *Queue) Dequeue() interface{} {
	return q.queue.PopTail()
}

func (q *Queue) Len() int {
	return q.queue.Len()
}

func NewQueue() *Queue {
	return &Queue{NewLinkedList()}
}
