package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	total := Sum(1, 2)
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}

func TestDivide(t *testing.T) {
	total, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide was incorrect, got: %d, want: %d.", total, 5)
	}

	if total != 5 {
		t.Errorf("Divide was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Errorf("Divide should return error when dividing by zero")
	}
}

func TestFailSum(t *testing.T) {
	total := Sum(1, 2)
	if total != 4 {
		t.Fail()
	}
	fmt.Println("Test Fail Sum Execution has been skipped")
}

func TestFailNowSum(t *testing.T) {
	total := Sum(1, 2)
	if total != 4 {
		t.FailNow()
	}
	fmt.Println("Test Fail Now Sum Execution has been skipped")
}

func TestFatalSum(t *testing.T) {
	total := Sum(1, 2)
	if total != 4 {
		t.Fatal("Test Fatal Sum Execution has been stopped")
	}
	fmt.Println("Test Fatal Sum Execution has been skipped")
}

func TestAssertSum(t *testing.T) {
	total := Sum(1, 2)
	assert.Equal(t, 3, total, "Sum should be 3")
	fmt.Println("Test Assert Sum Execution has been skipped")
}

func TestRequireSum(t *testing.T) {
	total := Sum(1, 2)
	require.Equal(t, 3, total, "Sum should be 3")
	fmt.Println("Test Require Sum Execution has been skipped")
}

func TestMassiveSum(t *testing.T) {
	for i := 0; i < 10; i++ {
		total := Sum(1, 2)
		assert.Equal(t, 3, total, "Sum should be 3")
	}
	fmt.Println("Test Massive Sum Execution has been skipped")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Test 1", 1, 2, 3},
		{"Test 2", 2, 3, 5},
		{"Test 3", 3, 4, 7},
		{"Test 4", 4, 5, 9},
		{"Test 5", 5, 6, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := Sum(tt.a, tt.b)
			assert.Equal(t, tt.expected, total, "Sum should be 3")
		})
	}
	fmt.Println("Test Table Sum Execution has been skipped")
}

func TestTable2Sum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errorMsg string
	}{
		{
			request:  Sum(1, 2),
			expected: 3,
			errorMsg: "Sum should be 3",
		},
		{
			request:  Sum(2, 3),
			expected: 5,
			errorMsg: "Sum should be 5",
		},
		{
			request:  Sum(3, 4),
			expected: 7,
			errorMsg: "Sum should be 7",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.request, tt.errorMsg)
		})
	}
	fmt.Println("Test Table 2 Sum Execution has been skipped")
}

// Positive Case : Login success
func TestLoginSuccess(t *testing.T) {
	result, err := Login("admin", "admin")
	assert.Nil(t, err)
	assert.True(t, result)
}

// Negative Case : Login failed
func TestLoginFailed(t *testing.T) {
	result, err := Login("admin", "admin1")
	if err == nil {
		t.Fail()
	}

	assert.False(t, result)
}
