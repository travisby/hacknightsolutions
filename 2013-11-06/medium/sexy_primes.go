package main

import "os"
import "strconv"
import "fmt"

func listSexy(n int) [][]int {
	fmt.Print(n)
	return [][]int{}
}

func main() {
	var n int64
	var _ error

	n, _ = strconv.ParseInt(os.Args[0], 10, 0)

	fmt.Print(listSexy(int(n)))
}
