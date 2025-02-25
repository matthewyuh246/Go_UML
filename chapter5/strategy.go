package main

import "fmt"

type SortStrategy interface {
	sort(list []int) []int
}

type BubbleSort struct{}

func (b BubbleSort) sort(list []int) []int {
	fmt.Println("BubbleSort")
	lenNumber := len(list)

	for i := 0; i < lenNumber; i++ {
		for j := 0; j < lenNumber-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	return list
}

type InsertionSort struct{}

func (i InsertionSort) sort(list []int) []int {
	fmt.Println("InsertionSort")
	lenNumber := len(list)

	for i := 1; i < lenNumber; i++ {
		tmp := list[i]
		j := i - 1
		for j >= 0 && list[j] > tmp {
			list[j] = list[j+1]
			j--
		}
		list[j+1] = tmp
	}
	return list
}

type SortContext struct {
	strategy SortStrategy
}

func (sc *SortContext) sort(list []int) []int {
	return sc.strategy.sort(list)
}

// func main() {
// 	list := []int{2, 5, 1, 8, 7, 3}
// 	b := BubbleSort{}
// 	context1 := SortContext{b}
// 	fmt.Println(context1.sort(list))
// 	i := InsertionSort{}
// 	context2 := SortContext{i}
// 	fmt.Println(context2.sort(list))

// }
