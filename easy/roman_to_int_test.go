package easy_test

import (
	"leetcode/easy"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	if easy.RomanToInt("III") != 3 {
		t.Error("error: 3")
	}
	if easy.RomanToInt("IV") != 4 {
		t.Error("error: 4")
	}

	if easy.RomanToInt("IX") != 9 {
		t.Error("error: 9")
	}

	if easy.RomanToInt("LVIII") != 58 {
		t.Error("error: 58")
	}
}
