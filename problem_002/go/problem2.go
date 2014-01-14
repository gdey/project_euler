package main

import (
	"fmt"
)

func fibonacci(quit chan (struct{})) <-chan (int) {
	result := make(chan (int), 1)
	go func() {
		l, c := 1, 1
		for {
			select {
			case <-quit:
				close(result)
				return
			case result <- c:
				l, c = c, l+c
			}
		}
	}()
	return result
}

func main() {
	quit := make(chan (struct{}))
	fib := fibonacci(quit)
	sum := 0
	for i := <-fib; i < 4000000; i = <-fib {
		if i%2 == 0 {
			sum += i
		}
	}
	close(quit)
	fmt.Println("Sum is", sum)
}
