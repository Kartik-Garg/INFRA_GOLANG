package main

import (
	"fmt"
	"time"
)

func PrintAlpha(start int, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(string('a' - 1 + i))
	}
}

func main() {
	go PrintAlpha(1, 5)
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from main")
}
