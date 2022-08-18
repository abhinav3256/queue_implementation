package main

import "fmt"

type Queue struct {
	Element []int
	Size    int
}

func main() {
	q := Queue{}
	q.Size = 4
	q.deque()
	q.enque(8)
	q.enque(81)
	q.enque(45)
	q.enque(63)
	q.deque()
	q.deque()
	fmt.Println(q.Element)
}
func (q *Queue) enque(x int) {

	if !q.isFull() {
		q.Element = append(q.Element, x)
	} else {
		fmt.Println("overflow")
	}

}

func (q *Queue) deque() {
	if !q.isEmpty() {
		q.Element = q.Element[1:]
	} else {
		fmt.Println("underflow")
	}

}

func (q *Queue) isEmpty() bool {
	if len(q.Element) == 0 {
		return true
	}
	return false
}

func (q *Queue) isFull() bool {
	if len(q.Element) == q.Size {
		return true
	} else {
		return false
	}
}
