package algorithms

func QuickSort(array []int) []int {

	return sortArray(array)
}

func sortArray(array []int) (sortedArray []int) {
	if len(array) <= 1 {
		return array
	}

	lessPivot := make([]int, 0, len(array))
	morePivot := make([]int, 0, len(array))
	equalPivot := make([]int, 0, len(array))

	pivot := getPivot(array)

	for _, v := range array {
		switch {
		case v < pivot:
			lessPivot = append(lessPivot, v)
		case v == pivot:
			equalPivot = append(equalPivot, v)
		case v > pivot:
			morePivot = append(morePivot, v)
		}
	}

	lessPivot = sortArray(lessPivot)
	morePivot = sortArray(morePivot)

	sortedArray = append(sortedArray, lessPivot...)
	sortedArray = append(sortedArray, equalPivot...)
	sortedArray = append(sortedArray, morePivot...)
	return
}

func getPivot(array []int) (pivot int) {
	middle := make([]int, 0, 3)
	for i, v := range array {
		if i <= (len(array)/2)+1 && i >= (len(array)/2)-1 {
			middle = append(middle, v)
		}
	}

	pivot = BubbleSort(middle)[1]

	return
}
