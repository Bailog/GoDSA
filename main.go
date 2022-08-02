package main

import (
	"Bailog/dsa/stack"
	"fmt"
)

func main() {
	// Init stack
	st := &stack.Stack{}
	st.Create(5)

	// Push into stack
	st.Push(22)
	st.Push(6)
	st.Push(9)
	st.Push(43)

	//Output
	fmt.Println("Stack:", st)

	// Pop
	st.Pop()

	//Output
	fmt.Println("Stack:", st)

	// Pop
	st.Pop()

	//Output
	fmt.Println("Stack:", st)

	// Peek
	fmt.Println("Stack:", st.Peek())

}
