package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x := 2
	y := 3

	result := Add(x, y)
	expected := 5

	assert.Equal(t, expected, result, "Add(2,3) should return 5")
}
