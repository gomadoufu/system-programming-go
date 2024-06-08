package concurrency

import (
	"fmt"
	"sync"
)

func waitGroup() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("First goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second goroutine")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All goroutines finished")
}
