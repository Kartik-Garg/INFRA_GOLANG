package main

import (
	"fmt"
	"sync"
)

func removeNums(s string, wg *sync.WaitGroup) {
	var words string
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char >= "a" && char <= "z" {
			words += char
		}
	}

	fmt.Print(words, " ")
	wg.Done()
}

func main() {
	arr := []string{"gopher123", "alpha99beta", "1cita2del3"}
	//have to remove the nums but using single goroutine
	var wg sync.WaitGroup
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		//call a go routine which will remove the nums
		go removeNums(arr[i], &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines are finished executing")

}

//reason only citadel is coming we didnt wait for enough time
