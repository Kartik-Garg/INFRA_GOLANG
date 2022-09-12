package main

import (
	"fmt"
)

func iterateOverArray(s []string, chanel1 chan []string) {
	var words []string
	//make empty channel here and get data from inner most function
	chanel2 := make(chan string)

	for _, k := range s {
		//send word and empty channel
		go removeNum(k, chanel2)
		//receive chanel2 here and add to array
		words = append(words, <-chanel2)

	}

	//send the words array back to outer func
	chanel1 <- words
}

func removeNum(s string, chanel2 chan string) {

	//we receive each word here
	var k string
	//iterating over word
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c >= "a" && c <= "z" {
			k += c
		}

	}
	//got one word, feed it to the internal (chanel2)
	chanel2 <- k
}

func main() {
	arr := []string{"gopher123", "alpha99beta", "1cita2del3"}

	chanel1 := make(chan []string)
	defer close(chanel1)

	//passed array and empty channel
	go iterateOverArray(arr, chanel1)
	//receive from the chanel1 and print
	fmt.Println(<-chanel1)

}
