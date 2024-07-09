package sumclice_test

import (
	"testing"
	"wb/sumclice"
)

func TestSumAsync(t *testing.T) {
	slice := sumclice.GenRandSlice(1000)
	expected := sumclice.SumSlice(slice)
	result := sumclice.SumSliceAsync(slice, 10)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func BenchmarkSumSync(b *testing.B) {
	numbers := sumclice.GenRandSlice(1000)
	for i := 0; i < b.N; i++ {
		sumclice.SumSlice(numbers)
	}
}

func BenchmarkSumAsync(b *testing.B) {
	numbers := sumclice.GenRandSlice(1000)
	for i := 0; i < b.N; i++ {
		sumclice.SumSliceAsync(numbers, 100)
	}
}
