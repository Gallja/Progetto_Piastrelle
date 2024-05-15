package main

type queue struct {
	head *queueNode
}

type queueNode struct {
	next  *queueNode
	value piastrella
}

func (q *queue) enqueue(value piastrella) {
	if q.head == nil {
		q.head = &queueNode{nil, value}
		return
	}

	node := q.head

	for node.next != nil {
		node = node.next
	}

	newNode := queueNode{nil, value}
	node.next = &newNode
}

func (q *queue) dequeue() piastrella {
	head := q.head
	q.head = q.head.next

	return head.value
}

func (q *queue) isEmpty() bool {
	return q.head == nil
}
