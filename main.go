package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {

	a := "foo"
	b := "foo"
	c := 4
	d := 5
	fmt.Println(isEqual(a, b))
	fmt.Println(isEqual(c, d))

	firstNums := []int32{1, 2, 3}
	secondNums := []uint32{4, 5, 6}
	thirdNums := []CustomInt{CustomInt(7), CustomInt(8), CustomInt(9)}
	fmt.Println(SumNumbers(firstNums))
	fmt.Println(SumNumbers(secondNums))
	fmt.Println(SumNumbers(thirdNums))

	arr := MyArray[int]{inner: []int{6, 4, 8, 9, 4, 0}}
	fmt.Println(arr.Max())

}

func isEqual[T comparable](a, b T) bool {
	return a == b
}

type Integer32 interface {
	~int32 | ~uint32
}

func SumNumbers[T Integer32](arr []T) T {
	var sum T
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

type CustomInt int32

// -----------------------------------------------------

type MyArray[T constraints.Ordered] struct {
	inner []T
}

func (m *MyArray[T]) Max() T {
	max := m.inner[0]
	for i := 0; i < len(m.inner); i++ {
		if m.inner[i] > max {
			max = m.inner[i]
		}
	}
	return max
}
