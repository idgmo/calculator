package main

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 1},
		{1000, 1000, 2000},
		{2, -2, 0},
		{0, -1, -1},
		{-1, 0, -1},
		{1, 0, 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := add(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func Test_divide(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1000, 1000, 1},
		{2, -2, -1},
		{0, -1, 0},
		{-1, 0, 0},
		{1, 0, 0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans, err := divide(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
				fmt.Printf("Error %s", err)
			}
		})
	}
}
