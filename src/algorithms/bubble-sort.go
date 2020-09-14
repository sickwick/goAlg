package algorithms

func BubbleSort(array []int) []int {
	for i := len(array); i > 0; i-- {
		for j := 1; j < i; j++ {
			if array[j] < array[j-1] {
				mid := array[j]
				array[j] = array[j-1]
				array[j-1] = mid
			}
		}
	}
	return array
}
