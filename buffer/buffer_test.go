package buffer

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCapacity(t *testing.T) {
	testCases := []struct {
		name             string
		capacity         uint
		expectedCapacity uint
	}{
		{"10", 10, 10},
		{"100", 100, 100},
		{"100_000", 100_000, 100_000},
		{"100_000_000", 100_000_000, 100_000_000},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			ringBuffer := NewBasicRingBuffer[int](testCase.capacity)
			assert.Equal(t, testCase.expectedCapacity, ringBuffer.GetCapacity())
		})

	}
}

// TestRead runs `runs` times and each time initializes
// a ring buffer of a random capacity and then runs
// a random set of times where it writes a value
// and subsequently reads it checking
// if the value is same
func TestRead(t *testing.T) {
	runs := 3
	maxBufferCapacity := 10
	maxTestTimes := 100
	maxValueToWrite := 10000

	for i := 0; i < runs; i++ {
		capacity := uint(rand.Intn(maxBufferCapacity))
		ringBuffer := NewBasicRingBuffer[int](capacity)

		testTimes := rand.Intn(maxTestTimes + 1)
		for j := 0; j < testTimes; j++ {
			valueToWrite := rand.Intn(maxValueToWrite)
			ringBuffer.Write(&valueToWrite)

			elementRead := ringBuffer.Read()
			assert.NotNil(t, elementRead)
			assert.Equal(t, valueToWrite, *elementRead)
		}
	}
}

func FuzzGetCapacity(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity uint) {
		ringBuffer := NewBasicRingBuffer[int](capacity)
		assert.Equal(t, capacity, ringBuffer.GetCapacity())
	})
}
