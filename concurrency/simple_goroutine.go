package concurrency

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("share by arguments:", c*c)
}

func simple() {
	go sub1(10)

	c := 20
	go func() {
		fmt.Println("share by closure:", c*c)
	}()
	time.Sleep(1 * time.Second)
}
