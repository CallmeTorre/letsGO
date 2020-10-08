package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(T *testing.T) {
	elements := []int{9, 8, 7, 6, 5}
	BubbleSort(elements)
	assert.NotNil(T, elements)
	assert.EqualValues(T, 5, len(elements))
	assert.EqualValues(T, 5, elements[0])
	assert.EqualValues(T, 6, elements[1])
	assert.EqualValues(T, 7, elements[2])
	assert.EqualValues(T, 8, elements[3])
	assert.EqualValues(T, 9, elements[4])

}

func TestBubbleSortBestCase(T *testing.T) {
	elements := []int{5, 6, 7, 8, 9}
	BubbleSort(elements)
	assert.NotNil(T, elements)
	assert.EqualValues(T, 5, len(elements))
	assert.EqualValues(T, 5, elements[0])
	assert.EqualValues(T, 6, elements[1])
	assert.EqualValues(T, 7, elements[2])
	assert.EqualValues(T, 8, elements[3])
	assert.EqualValues(T, 9, elements[4])

}

func TestBubbleSortNilSlice(T *testing.T) {
	BubbleSort(nil)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(T *testing.T) {
	elements := getElements(5)
	assert.NotNil(T, elements)
	assert.EqualValues(T, 5, len(elements))
	assert.EqualValues(T, 4, elements[0])
	assert.EqualValues(T, 3, elements[1])
	assert.EqualValues(T, 2, elements[2])
	assert.EqualValues(T, 1, elements[3])
	assert.EqualValues(T, 0, elements[4])

}

func BenchmarkBubbleSort10(B *testing.B) {
	elements := getElements(10)
	for i := 0; i < B.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort1000(B *testing.B) {
	elements := getElements(1000)
	for i := 0; i < B.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort10(B *testing.B) {
	elements := getElements(10)
	for i := 0; i < B.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkSort1000(B *testing.B) {
	elements := getElements(10)
	for i := 0; i < B.N; i++ {
		sort.Ints(elements)
	}
}
