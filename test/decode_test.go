package main

import (
	"blog-api/blog"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	cases := map[string]int{
		"12":  2,
		"226": 3,
		"0":   0,
		"10":  1,
		"100": 0,
		"101": 1,
	}

	for input, expected := range cases {
		got := blog.NumDecodings(input)
		if got != expected {
			t.Errorf("NumDecodings(%s) = %d; want %d", input, got, expected)
		}
	}
}
