package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func printSomethingExtended(s string, wg *sync.WaitGroup) { // We create this function to extend the functionality with WaitGroups
	defer wg.Done() // Decrement wg by 1. With 'defer' we specify to do it at the very end.
	fmt.Println(s)
}

func main() {
	go printSomething("This is a thing to be printed.")
	go printSomething("This is another string.")
	printSomething("This is the third string.")
	time.Sleep(2 * time.Second) // Wait to allow goroutines to finish, if not, program may exit before they run.

	var wg sync.WaitGroup // WaitGroup helps us to wait for things to finish.
	words := []string{"apple", "banana", "cherry", "alpha"}
	wg.Add(4) // We pass a 4 since we need to wait for the 4 elements to be printed. Also could use len(words)

	for i, word := range words {
		go printSomethingExtended(fmt.Sprintf("%d: %s", i, word), &wg)
	}
	wg.Wait() // We indicate to wait for waitgroups to be '0'
	// This may be returned in different order everytime we run it, since the threads can terminate at different time.
}
