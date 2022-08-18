package main

import "fmt"

type Queue struct {
	Element []int
	Size    int
}

func main() {
	q := Queue{}
	q.Size = 5

	q.enque(25)
	q.enque(12)
	q.enque(17)
	q.enque(87)
	q.enque(26)
	q.enque(38)

	fmt.Println(q.Element)
	q.deque()
	fmt.Println(q.Element)

}

func (q *Queue) enque(element int) int {
	if q.getLength() == q.Size {
		fmt.Println("queue overflow")
		return 0
	}
	q.Element = append(q.Element, element)
	fmt.Println("Elemnt added:", element)
	return 1
}

func (q *Queue) deque() int {
	// if q.isEmpty() {
	// 	fmt.Println("queue underflow")
	// 	return 0
	// }
	q.Element = q.Element[1:]
	return 1
}

func (q *Queue) getLength() int {

	return len(q.Element)
}

func (q *Queue) isEmpty() bool {

	if len(q.Element) > 0 {
		return false
	} else {
		return true
	}
}
