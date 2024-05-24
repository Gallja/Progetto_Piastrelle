package main

type queue struct {
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	next  *queueNode
	value piastrella
}

func (q *queue) enqueue(value piastrella) {
	if q.head == nil {
		q.head = &queueNode{nil, value}
		q.tail = q.head
		return
	}

	newNode := &queueNode{nil, value}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *queue) dequeue() piastrella {
	head := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return head.value
}

func (q *queue) isEmpty() bool {
	return q.head == nil
}
