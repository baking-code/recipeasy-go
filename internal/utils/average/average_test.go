package average

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	avg := NewAverageNumberCounter()
	avg.AddNewNumber(1)
	avg.AddNewNumber(2)
	avg.AddNewNumber(3)
	avg.AddNewNumber(4)
	assert.Equal(t, 2.5, avg.Average)
	avg.AddNewNumber(5)
	assert.Equal(t, 3.0, avg.Average)

	avg.Reset()
	avg.AddNewNumber(1)
	avg.AddNewNumber(2)
	avg.AddNewNumber(-3)
	avg.AddNewNumber(4)
	assert.Equal(t, 1.0, avg.Average)
	avg.AddNewNumber(6)
	assert.Equal(t, 2.0, avg.Average)

	avg.Reset()
	avg.AddNewNumber(0)
	avg.AddNewNumber(0)
	avg.AddNewNumber(0)
	assert.Equal(t, 0.0, avg.Average)
}

func TestZeroAverage(t *testing.T) {
	avg := NewAverageNumberCounter()

	avg.AddNewNumber(0)
	assert.Equal(t, 0.0, avg.Average)
}

const float64EqualityThreshold = 1e-9

func FuzzTestAverage(f *testing.F) {
	testcases := []int{1, 2, 3, 4, 5}
	var almostEqual = func(a, b float64) bool {
		return math.Abs(a-b) <= float64EqualityThreshold
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	avg := NewAverageNumberCounter()
	count := 0
	sum := 0
	f.Fuzz(func(t *testing.T, in int) {
		count++
		sum += in
		fmt.Println(in)
		avg.AddNewNumber(int(in))
		expected := float64(sum) / float64(count)

		result := avg.Average
		if !almostEqual(result, expected) {
			t.Errorf("Expected: %v, got: %v", expected, result)
		}
	})

}

func TestStore(t *testing.T) {
	avg := NewAverageNumberCounter()
	avg.AddNewNumber(1)
	avg.AddNewNumber(2)
	avg.AddNewNumber(3)
	avg.AddNewNumber(4)

	avg.AddNewNumber(4)
	avg.AddNewNumber(-3)
	avg.AddNewNumber(4)

	assert.Equal(t, 3, avg.Store[4])
	assert.Equal(t, 1, avg.Store[3])
	assert.Equal(t, 0, avg.Store[5])

}
