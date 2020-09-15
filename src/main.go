package main

import (
	"algorithms"
	"fmt"
	"time"
)

func main() {
	arr := []int{4, 12, 3, 2, 23, 2, 5, 0, 90, 8, 4, 65}
	start := time.Now()
	fmt.Println(algorithms.BubbleSort(arr))
	//fmt.Println(algorithms.QuickSort(arr))
	fmt.Println(time.Since(start).Microseconds())
}
