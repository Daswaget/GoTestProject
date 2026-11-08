package operations

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	var testCases = []struct {
		x, y, want float64
	}{
		{0, 0, 0},
		{2, 2, 4},
		{-2, 4, 2},
		{0.5, 2, 2.5},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := Addition(tc.x, tc.y)
			if result != tc.want {
				t.Errorf("Got %f, want %f", result, tc.want)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	var testCases = []struct {
		x, y, want float64
	}{
		{0, 0, 0},
		{2, 2, 0},
		{-2, 4, -6},
		{0.5, 2, -1.5},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := Subtraction(tc.x, tc.y)
			if result != tc.want {
				t.Errorf("Got %f, want %f", result, tc.want)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	var testCases = []struct {
		x, y, want float64
	}{
		{0, 0, 0},
		{2, 2, 4},
		{-2, 4, -8},
		{0.5, 2, 1},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := Multiplication(tc.x, tc.y)
			if result != tc.want {
				t.Errorf("Got %f, want %f", result, tc.want)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	var testCases = []struct {
		x, y, want float64
	}{
		{2, 2, 1},
		{-2, 4, -0.5},
		{0.5, 2, 0.25},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := Division(tc.x, tc.y)
			if result != tc.want {
				t.Errorf("Got %f, want %f", result, tc.want)
			}
		})
	}
}
