package main

import (
	"algorithms"
	"fmt"
)

func main() {
	arr := []int{4, 12, 3, 2, 23, 2, 5, 0, 90, 8, 4, 65}
	fmt.Println(algorithms.BubbleSort(arr))
}
