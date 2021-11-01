package easy_test

import (
	"leetcode/easy"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	if easy.BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Error("error: 4")
	}

	if easy.BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		t.Error("error: 2")
	}
}
