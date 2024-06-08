package concurrency

import "fmt"

func Concurrency() {
	fmt.Println("Concurrency")
	simple()
	mutex()
	waitGroup()
}
