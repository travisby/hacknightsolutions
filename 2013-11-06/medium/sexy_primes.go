package main

import "os"
import "strconv"
import "fmt"

func isDivisibleBy(n int, m int) bool {
	return n%m == 0
}

func divisors(n int) []int {
	var results []int

	for i := 1; i <= n; i++ {
		if isDivisibleBy(n, i) {
			results = append(results, i)
		}
	}

	return results
}

func isPrime(n int) bool {
	return len(divisors(n)) == 2
}

func listSexy(n int) [][]int {
	var results [][]int

	for i := 0; i <= n; i++ {
		m := i + 6
		if isPrime(m) && isPrime(i) {
			results = append(results, []int{i, m})
		}
	}

	return results
}

func main() {
	var n int
	var _ error

	n, _ = strconv.Atoi(os.Args[1])

	fmt.Print(listSexy(n))
	fmt.Print("\n")
}
