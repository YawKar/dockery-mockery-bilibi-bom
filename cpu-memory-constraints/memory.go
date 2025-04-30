package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	megabytes := flag.Arg(0)
	if megabytes == "" {
		fmt.Fprintln(os.Stderr, "you should've passed megabytes!")
		os.Exit(1)
	}
	mb, err := strconv.Atoi(megabytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "you should have passed a positive number: %s\n", err.Error())
		os.Exit(1)
	}

	data := make([]byte, mb*1024*1024)
	for i := range data {
		data[i] = byte(i)
	}
	fmt.Println(len(data))
	return
}
