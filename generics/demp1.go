package main

import "fmt"

func main() {

	var s = stack[int]{value: []int{1, 2, 3, 4}}
	fmt.Println(s.sum_stack())
	s1 := stack[int]{[]int{1, 2, 3, 4}}
	fmt.Println(s1.findOne(1))
	s2 := stack[float64]{[]float64{1.2, 2.3, 3.4, 4.5}}
	fmt.Println(s2.findOne(1.3))
}

type stack[T int | float64] struct {
	value []T
}

func (s stack[T]) sum_stack() T {
	var sum T
	for _, t := range s.value {
		sum += t
	}
	return sum
}

func (s stack[T]) findOne(tt T) bool {
	for _, t := range s.value {
		if t == tt {
			return true
		}
	}
	return false
}
