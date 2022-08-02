package queue

type Queue struct {
	front int
	rear  int
	items []int
}

func (q *Queue) Create(max int) []int {
	q.items = make([]int, 0, max)
	q.front, q.rear = -1, -1
	return q.items
}

func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

func (q *Queue) IsFull() bool {
	return q.front == 0 && q.rear == cap(q.items)-1
}

func (q *Queue) Enqueue(val int) int {
	if q.IsFull() {
		return -1
	} else {
		if q.rear == -1 {
			q.front = 0
		}
		q.rear += 1
		q.items = append(q.items, val)
	}
	return -1
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	} else {
		elem := q.items[q.front]
		if q.front >= q.rear {
			q.front, q.rear = -1, -1
		} else {
			q.front += 1
		}
		return elem
	}
}

func (q *Queue) Show() []int {
	return q.items[q.front:]
}
