package main

import (
	"fmt"
)


func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n<= 3 {
		return true
	}
	for i:=2; i<=2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func printPrimeFactors(n int) {
	if isPrime(n) {
		fmt.Println(n)
		return
	}
	primes := []int{}
	for i:=2; i<=n/2; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes)
	factors := []int{}
	for i:= 0; i<len(primes); i++ {
		prime := primes[i]
		if n%prime == 0 {
			n = n/prime
			factors = append(factors, prime)
			if n % prime == 0 {
				i--
			}
		}
	}
	for _, factor := range factors {
		fmt.Println(factor)
	}

}

func main() {
	var n int
	fmt.Println("n: ")
	fmt.Scanf("%d", &n)
	if isPrime(n) {
		fmt.Println("n is Prime")
	} else {
		fmt.Println("n is not Prime")
	}
	printPrimeFactors(n)
}
