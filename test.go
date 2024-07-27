package main

import (
	"fmt"
	"sync"
)

// Подсчитать сумму квадратов чисел от 1 до N. Каждая операция возведения в квадрат должна происходить в отдельной горутине.
func firstTest() {
	n := quad(3)
	fmt.Println(n)
	cn := cuncurentQuad(3)
	fmt.Println(cn)
}

func cuncurentQuad(n int) int {
	sum := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 1; i <= n; i++ {
		wg.Add(1)
		i := i
		go func() {
			mu.Lock()
			defer mu.Unlock()
			sum += i * i
			wg.Done()
		}()
	}
	wg.Wait()
	return sum
}

func quad(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}
