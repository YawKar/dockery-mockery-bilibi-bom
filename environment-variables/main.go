package main

import (
	"fmt"
	"os"
)

func main() {
	color, ok := os.LookupEnv("COLOR")
	if !ok {
		fmt.Fprintln(os.Stderr, "no COLOR was set!")
		os.Exit(1)
	}
	name, ok := os.LookupEnv("NAME")
	if !ok {
		fmt.Fprintln(os.Stderr, "no NAME was set!")
		os.Exit(1)
	}
	fmt.Printf("Color: %s; Name: %s;\n", color, name)
}
