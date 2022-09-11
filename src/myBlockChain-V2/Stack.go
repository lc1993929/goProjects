package main

import "container/list"

type Stack struct {
	data *list.List
	size int
}

func NewStack() *Stack {
	q := new(Stack)
	q.init()
	return q
}

func (q *Stack) init() {
	q.data = list.New()
}

func (q *Stack) Size() int {
	return q.size
}

func (q *Stack) Empty() bool {
	return q.size == 0
}

func (q *Stack) Top() interface{} {
	return q.data.Back().Value
}

func (q *Stack) Push(value interface{}) {
	q.data.PushBack(value)
	q.size++
}

func (q *Stack) Pop() interface{} {
	if q.size > 0 {
		tmp := q.data.Back()
		q.data.Remove(tmp)
		q.size--
		return tmp.Value
	}
	return nil
}
