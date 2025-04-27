package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	for {
		t := <-ticker.C
		fmt.Println(t)
	}
}
