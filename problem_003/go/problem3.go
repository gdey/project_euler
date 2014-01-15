package main

import (
	"flag"
	"fmt"
	"math"
)

var number = flag.Int("number", 600851475143, "Which number do you want the prime factors of; defaults to 600851475143")

func getNextPrime(sn int) int {
	i := 3
	for ; i <= int(math.Sqrt(float64(sn))); i++ {
		if (sn % i) == 0 {
			sn = sn + 2 // Only look at the odd numbers.
			i = 2
		}
	}
	return sn
}

func genPrimes(quitChan <-chan (struct{})) <-chan (int) {
	resultChan := make(chan (int), 2)
	go func() {
		resultChan <- 2
		sn := 3
		resultChan <- sn

		for {
			select {
			case <-quitChan:
				return
			default:
				sn = getNextPrime(sn + 2) // Only look at the odd numbers.
				resultChan <- sn
			}
		}
	}()
	return resultChan
}

func main() {
	flag.Parse()
	quit := make(chan (struct{}))
	primes := genPrimes(quit)
	factors := make([]int, 0, 10)

	num := *number
	for num > 1 {
		p := <-primes
		for num%p != 0 {
			p = <-primes
		}
		factors = append(factors, p)
		num = num / p
	}
	fmt.Println("For the number", *number, "prime factors are", factors)
	fmt.Println("The largest factor is", factors[len(factors)-1])
	close(quit)
}
