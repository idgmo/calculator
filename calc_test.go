package main

import (
	"fmt"
	"testing"
)

func Test_sum(t *testing.T) {
	tests := []struct {
		a    string
		want int
	}{

		{"", 0},
		{"3", 3},
		{"1,2", 3},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := sum(tt.a)
			if ans == 0 {
				t.Log("ans is 0")
			}
			t.Log(ans)
		})
	}
}
