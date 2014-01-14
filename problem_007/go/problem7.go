package main

import (
	"fmt"
	"math"
	"flag"
)

var whichPrime = flag.Int("whichPrime", 10001, "Which prime number do you want; defaults to 10001")

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
	for i := 0; i < *whichPrime - 1 ; i++ {
		_ = <-primes // eat up the other primes
	}
	fmt.Println("The",*whichPrime,"prime is", <-primes)
	close(quit)
}
