package main

import (
	"fmt"
	"github.com/mrNobody95/Gate/strategies"
)

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println("arr2 1:", arr2)
	fmt.Println("arr1 1:", arr1)
	arr2[1] = 9
	fmt.Println("arr2 2:", arr2)
	fmt.Println("arr1 2:", arr1)
}

func defaultStrategy() {
	strategy := strategies.Strategy{}

	strategy.Validate()
	strategy.CollectPrimaryData()
}
