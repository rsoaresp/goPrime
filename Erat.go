package Erat

import "math"

// Implementation of the sieve of Eratosthenes
func Erat(n int) []bool {

	// We begin with all numbers marked as primes,
	// since later we will delete all the multiples
	// of previously know primes, starting from 2.
	x := make([]bool, n)
	for i := 0; i < n; i++ {
		x[i] = true
	}

	// This is the Sieve of Eratosthenes algorithm. We have
	// an array with all numbers from 2 to n marked as true.
	// Beginning with 2, that is a prime number, we will remove
	// all of its multiples from the list, marking them as false.
	// Then we move to the next integer in the list not false and
	// repeate the steps.
	var largestNum int = int(math.Sqrt(float64(n)))
	for p := 2; p < largestNum; p++ {
		if x[p] == true {
			start := int(math.Pow(float64(p), 2))
			for j := start; j < n; j += p {
				x[j] = false
			}
		}
	}

	return x
}
