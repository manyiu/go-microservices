package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	element := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	actual := BubbleSort(element)

	assert.NotNil(t, actual)
	assert.EqualValues(t, 5, len(actual))
	assert.EqualValues(t, expected, actual)
}

func TestBubbleSortBestCase(t *testing.T) {
	element := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	actual := BubbleSort(element)

	assert.NotNil(t, actual)
	assert.EqualValues(t, 5, len(actual))
	assert.EqualValues(t, expected, actual)
}

func TestBubbleSortWithEmptyElement(t *testing.T) {
	element := []int{}
	expected := []int{}

	actual := BubbleSort(element)

	assert.NotNil(t, actual)
	assert.EqualValues(t, 0, len(actual))
	assert.EqualValues(t, expected, actual)
}

func getElement(n int) []int {
	result := make([]int, n)

	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func TestGetElements(t *testing.T) {
	els := getElement(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, []int{4, 3, 2, 1, 0}, els)
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElement(10)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElement(1000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElement(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElement(100000)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
