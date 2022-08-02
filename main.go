package main

import (
	"Bailog/dsa/queue"
	"Bailog/dsa/stack"
	"fmt"
)

func stackExample() {
	// Init stack
	st := &stack.Stack{}
	st.Create(10)
	defer func() {
		st = nil
	}()

	// Push into stack
	st.Push(22)
	st.Push(43)
	st.Push(43)
	st.Push(43)
	//Output
	fmt.Println("Push(4) Stack:", st)

	// Pop
	st.Pop()
	st.Pop()
	//Output
	fmt.Println("Pop(2) Stack:", st)

	st.Push(22)
	st.Push(232)
	st.Push(24)
	st.Push(25)
	st.Push(26)
	//Output
	fmt.Println("Push(5) Stack:", st)

	// Peek
	fmt.Println("Peek Stack:", st.Peek())
}

func queueExample() {
	//Init queue
	qu := &queue.Queue{}
	qu.Create(5)
	fmt.Println("Queue:", qu)

	//add into queue
	qu.Enqueue(10)
	qu.Enqueue(11)
	qu.Enqueue(12)
	fmt.Println("Enqueue Queue:", qu)

	//remove from queue
	qu.Dequeue()
	fmt.Println("Dequeue Queue:", qu)

	qu.Enqueue(13)
	//Show
	fmt.Println("Enqueue Queue:", qu.Show())
}

func main() {
	// Init stack and do all provided methods
	stackExample()

	// Init queue and do all provided methods
	queueExample()
}
