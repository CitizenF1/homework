package main

import (
	"fmt"
	"time"
	"wb/sumclice"
)

// Написать функцию выполняющую батчинг значений из канала c и возвращающую канал батчей, пополняемых асинхронного из канала c
func main() {
	// c := make(chan any, 10)
	// cOut := doBatching(c, 2)
	// go func() {
	// 	for i := range cOut {
	// 		fmt.Println(i) // [1,2] [3,4]
	// 	}
	// }()

	// c := make(chan any, 10)
	// cOut := doBatchingTime(c, 2, time.Nanosecond*1)
	// go func() {
	// 	for i := range cOut {
	// 		fmt.Println(i) // [1,2] [3,4]
	// 	}
	// }()
	// c <- 1
	// c <- 2
	// c <- 3
	// c <- 4
	// c <- 5
	// close(c)
	// time.Sleep(time.Second * 1)

	// ===================================

	c := make(chan any)
	timeout := 2 * time.Millisecond
	cOut := doBatchingTime(c, 2, timeout)

	go func() {
		for batch := range cOut {
			fmt.Printf("Batch: %v \n", batch)
		}
	}()

	for i := 1; i < 5; i++ {

		c <- i
	}
	c <- 22
	close(c)

	time.Sleep(time.Second * 1)
	fmt.Println("=======================")
	// ===================================

	slice := sumclice.GenRandSlice(10)

	sum := sumclice.SumSliceAsync(slice, 2)

	fmt.Println("SUM: ", sum)

	sumSync := sumclice.SumSlice(slice)

	fmt.Println("SUM Sync: ", sumSync)

}

func doBatching(c chan any, batchSize int) chan []any {
	batches := make(chan []any, batchSize)
	values := make([]any, 0, batchSize)

	go func() {
		for i := range c {

			values = append(values, i)

			if len(values) == batchSize {
				batches <- values
				values = nil
			}

		}

		if len(values) > 0 {
			batches <- values
		}

		fmt.Println("hello")
	}()

	return batches
}

// home
// timeTiker
// select
func doBatchingTime(c chan any, batchSize int, timeout time.Duration) chan []any {
	batches := make(chan []any)
	values := make([]any, 0, batchSize)

	go func() {
		timer := time.NewTimer(timeout)

		for {
			select {
			case i, ok := <-c:

				if !ok {
					if len(values) > 0 {
						batches <- values
					}
					return
				}

				values = append(values, i)

				if len(values) == batchSize {
					batches <- values
					values = make([]any, 0, batchSize)
					if !timer.Stop() {
						<-timer.C
					}
					timer.Reset(timeout)
				}

			case <-timer.C:
				if len(values) > 0 {
					batches <- values
					values = make([]any, 0, batchSize)
				}
				timer.Reset(timeout)

			case <-time.After(timeout):
				if len(values) > 0 {
					batches <- values
				}
				close(batches)
				return
			}

		}
	}()

	return batches
}
