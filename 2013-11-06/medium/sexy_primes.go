package main

import "os"
import "strconv"
import "fmt"

func isDivisibleBy(n int, m int) bool {
	return n%m == 0
}

func divisors(n int) []int {
	var recur func(int, int, []int) []int
	recur = func(a int, b int, r []int) []int {
		if b < 1 {
			return r
		} else if isDivisibleBy(a, b) {
			r = append(r, b)
		}

		return recur(a, b-1, r)
	}
	return recur(n, n, make([]int, 0))
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
