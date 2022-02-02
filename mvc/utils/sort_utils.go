package utils

import "sort"

func BubbleSort(element []int) []int {
	for i := 0; i < len(element); i++ {
		for j := 0; j < len(element)-1; j++ {
			if element[j] > element[j+1] {
				element[j], element[j+1] = element[j+1], element[j]
			}
		}
	}

	return element
}

func Sort(els []int) []int {
	if len(els) < 5000 {
		return BubbleSort(els)
	}

	sort.Ints(els)
	return els
}
