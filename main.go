package main

import "github.com/rauljordan/long-running-ci-poc/bold/staker"

func main() {
	stkr := staker.New()
	stkr.Foo()
}

func add(x int, y int) int {
	return x + y
}
