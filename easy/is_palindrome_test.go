package easy_test

import (
	"leetcode/easy"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !easy.IsPalindrome(121) {
		t.Error("error: 121")
	}

	if easy.IsPalindrome(-121) {
		t.Error("error: -121")
	}

	if easy.IsPalindrome(1232) {
		t.Error("error: 1232")
	}
}
