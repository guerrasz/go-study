package main

import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func BenchmarkAdd(b *testing.B) {
	for b.Loop() {
		Add(2, 3)
	}
}

func TestAddSubtests(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{a: 2, b: 3, expected: 5},
		{a: -1, b: 1, expected: 0},
		{a: 0, b: 0, expected: 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{a: 2, b: 3, expected: 5},
		{a: -1, b: 1, expected: 0},
		{a: 0, b: 0, expected: 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
