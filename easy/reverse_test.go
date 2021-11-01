package easy_test

import (
	"leetcode/easy"
	"testing"
)

func TestReverse(t *testing.T) {
	if easy.Reverse(1231) != 1321 {
		t.Error("error: 1231")
	}
	if easy.Reverse(-1231) != -1321 {
		t.Error("error: -1231")
	}
	if easy.Reverse(0) != 0 {
		t.Error("error: 0")
	}
	if easy.Reverse(120) != 21 {
		t.Error("error: 120")
	}
	// if easy.Reverse(1534236469) != 9646324351 {
	// 	t.Error("error: 1534236469")
	// }
}
