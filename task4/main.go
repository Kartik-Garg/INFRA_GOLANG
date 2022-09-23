package main

import (
	"fmt"
)

func iterateOverArray(s []string, stringArrayChannel chan []string) {
	var words []string
	//make empty channel here and get data from inner most function
	wordsChannel := make(chan string)

	for _, k := range s {
		//send word and empty channel
		go removeNum(k, wordsChannel)
		//receive wordsChannel here and add to array
	}

	for i := 0; i < len(s); i++ {
		words = append(words, <-wordsChannel)
	}
	//send the words array back to outer func
	stringArrayChannel <- words
}

func removeNum(s string, wordsChannel chan string) {

	//we receive each word here
	var k string
	//iterating over word
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c >= "a" && c <= "z" {
			k += c
		}

	}
	//got one word, feed it to the internal (wordsChannel)
	wordsChannel <- k
}

func main() {
	arr := []string{"gopher123", "alpha99beta", "1cita2del3"}

	stringArrayChannel := make(chan []string)
	defer close(stringArrayChannel)

	//passed array and empty channel
	go iterateOverArray(arr, stringArrayChannel)
	//receive from the stringArrayChannel and print
	fmt.Println(<-stringArrayChannel)

}
