package queue

import "fmt"

func ExampleQueue_Push() {
	q:=Queue{1}
	q.Push(1)
	q.Push("demon")
	q.Push(true)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	// output:
	// 1
	// 1
	// demon
	// true

}
