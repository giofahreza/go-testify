package helper

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	result, _ := Add(1, 2)
	if result != 4 {
		t.Error("Expected 3, got ", result)
	}
}

func TestAddFail(t *testing.T) {
	_, err := Add(1, -10)
	if err != errors.New("Negative numbers are not allowed") {
		t.Fail()
	}

	fmt.Println("Add(1, -10) failed as expected")
}

func TestAddFailNow(t *testing.T) {
	_, err := Add(1, -10)
	if err != errors.New("Negative numbers are not allowed") {
		t.FailNow()
	}

	fmt.Println("Add(1, -10) failed as expected")
}

func TestSomethingAssert(t *testing.T) {
	a := true
	b := true
	assert.Equal(t, a, b, "The result success")
}

func TestSomethingRequire(t *testing.T) {
	a := true
	b := true
	require.Equal(t, a, b, "The result success")
}

func TestSomethingAssertLess(t *testing.T) {
	a := 1
	b := 2
	assert.Less(t, a, b, "The result success")
}

func TableTestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Test 1", 1, 2, 4},
		{"Test 2", 1, -10, 0},
		{"Test 3", 1, 10, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAdd1(t *testing.T) {
	result, _ := Add(1, 2)
	assert.Equal(t, 4, result)
}

func TestAdd2(t *testing.T) {
	result, _ := Add(1, -10)
	assert.Equal(t, 0, result)
}

func TestAdd3(t *testing.T) {
	result, _ := Add(1, 10)
	assert.Equal(t, 12, result)
}
