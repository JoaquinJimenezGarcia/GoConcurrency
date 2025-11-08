package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	msg = s
	printMessage()
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	messages := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}
	wg.Add(len(messages))

	for _, message := range messages {
		go updateMessage(message, &wg)
	}

	wg.Wait()
}
