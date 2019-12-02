/**
 * @Author: DollarKillerX
 * @Description: 队列实现
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午1:57 2019/11/26
 */
package queue

// 后面性能测试这个队列比较消耗cpu就弃用了

import (
	"sync"
)

type queueNode struct {
	Next *queueNode
	Data interface{}
}

// Queue implements a FIFO data structure.
type Queue struct {
	sync.Mutex
	size       int
	head, tail *queueNode
}

// Append adds the data to the end of the Queue.
func (q *Queue) Append(data interface{}) {
	element := new(queueNode)

	q.Lock()
	defer q.Unlock()

	q.size++
	if q.head == nil {
		q.head = element
	}

	end := q.tail
	if end != nil {
		end.Next = element
	}
	q.tail = element
	element.Data = data
}

// Next returns the data at the front of the Queue.
func (q *Queue) Next() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()

	if q.head == nil {
		return nil, false
	}

	q.size--
	element := q.head
	q.head = element.Next
	if q.tail == element {
		q.tail = nil
	}
	element.Next = nil
	return element.Data, true
}

// Empty returns true if the Queue is empty.
func (q *Queue) Empty() bool {
	q.Lock()
	defer q.Unlock()

	if q.head == nil {
		return true
	}
	return false
}

// Len returns the current length of the Queue
func (q *Queue) Len() int {
	return q.size
}
