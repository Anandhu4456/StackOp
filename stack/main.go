package main

import "fmt"

type Stack struct {
	datas []int
}

// creating push method.......

func (s *Stack) Push(val int) {
	s.datas = append(s.datas, val)
}

// creating pop method.........

var popVal int

func (s *Stack) Pop() int {
	popVal = s.datas[len(s.datas)-1]
	s.datas = s.datas[:len(s.datas)-1]
	return popVal
}

// searching....

func (s *Stack) Search(val int) int {
	for index, value := range s.datas {
		if value == val {
			fmt.Printf("%v found at position %v: ", value, index+1)
		}
	}
	fmt.Println()
	return val

}

// peek....

func (s *Stack) Peek() {
	lastVal := s.datas[len(s.datas)-1]
	fmt.Println("The last entered value is : ", lastVal)
}

func main() {
	stackList := &Stack{}

	stackList.Push(200)
	stackList.Push(300)
	stackList.Push(500)
	stackList.Push(600)
	stackList.Push(800)
	stackList.Push(1000)
	fmt.Println(stackList)

	// stackList.Pop()
	// fmt.Println("The popped value is : ",popVal)
	fmt.Println(stackList)
	stackList.Search(300)
	stackList.Peek()
}
