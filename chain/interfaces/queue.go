package interfaces

import (
	"container/heap"
	"encoding/hex"
	"encoding/json"
)

type QueueElement struct {
	value string
	priority int
	index int
}

type Queue []*QueueElement

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return q[i].priority > q[j].priority
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *Queue) Push(x interface{}) {
	n := len(*q)
	queueElement := x.(*QueueElement)
	queueElement.index = n
	*q = append(*q, queueElement)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	queueElement := old[n-1]
	old[n-1] = nil  // avoid memory leak
	queueElement.index = -1 // for safety
	*q = old[0 : n-1]
	return queueElement
}

// update modifies the priority and value of an QueueElement in the queue.
func (q *Queue) update(queueElement *QueueElement, value string, priority int) {
	queueElement.value = value
	queueElement.priority = priority
	heap.Fix(q, queueElement.index)
}

func NewBlockQueue(items []*Block) Queue {

	q := make(Queue, len(items))
	i := 0

	for value, priority := range items {
		val, _ := json.Marshal(value)

		q[i] = &QueueElement{
			value:    hex.EncodeToString(val),
			priority: int(priority.Index),
			index:    i,
		}

		i++
	}
	return q
}

func AddToBlockQueue(item *Block, q Queue) Queue {

	val, _ := json.Marshal(item)
	it := &QueueElement{
		value: hex.EncodeToString(val),
		priority: int(item.Index),
		index: q.Len()-1,
	}
	q.Push(it)

	return q
}