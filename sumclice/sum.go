package sumclice

import (
	"sync"

	"golang.org/x/exp/rand"
)

func GenRandSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(10) + 1
	}
	return slice
}

func SumSliceAsync(slice []int, size int) int {
	var wg sync.WaitGroup
	ch := make(chan int, len(slice)/size)
	sum := 0

	for i := 0; i < len(slice); i += size {
		wg.Add(1)
		e := i + size

		if e > len(slice) {
			e = len(slice)
		}

		// fmt.Println("Slice: ", slice[i:e])

		go func(batch []int) {
			defer wg.Done()
			sum := 0
			for _, v := range batch {
				sum += v
			}
			ch <- sum

		}(slice[i:e])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for batchSum := range ch {
		sum += batchSum
	}

	return sum
}

func SumSlice(slice []int) int {
	sum := 0
	for i := range slice {
		sum += slice[i]
	}
	return sum
}
