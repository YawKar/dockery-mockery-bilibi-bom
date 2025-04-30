package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func parallelFibonacci(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		if n <= 2 {
			out <- 1
			return
		}
		c1 := parallelFibonacci(n - 1)
		c2 := parallelFibonacci(n - 2)
		out <- (<-c1 + <-c2)
	}()
	return out
}

func main() {
	flag.Parse()
	nStr := flag.Arg(0)
	if nStr == "" {
		fmt.Fprintln(os.Stderr, "you should've pass a positive number!")
		os.Exit(1)
	}
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "you should've pass a positive number! error: %s\n", err.Error())
		os.Exit(1)
	}
	if n <= 0 {
		fmt.Fprintf(os.Stderr, "you should've pass a positive number! error: %d is <= 0\n", n)
		os.Exit(1)
	}

	start := time.Now()
	resCh := parallelFibonacci(n)
	fmt.Println("Result:", <-resCh)
	duration := time.Since(start)
	fmt.Printf("Took: %s\n", duration)
}
