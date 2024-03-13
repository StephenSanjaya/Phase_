package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	data := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{17, "17"},
	}

	for _, d := range data {
		result := FizzBuzz(d.input)
		if result != d.expected {
			t.Errorf("Input %d, expected %s, but got %s", d.input, d.expected, result)
		}
	}
}
