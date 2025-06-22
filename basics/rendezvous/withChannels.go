package main

import (
	"fmt"
	"sync"
)

// A1 -> B2
// A1 should be done before B2
// B1 -> A2
// B1 should be done before A
func main() {
	// Note: If we use unbuffered channel, it will be deadlock
	a1Done := make(chan struct{}, 1)
	b1Done := make(chan struct{}, 1)
	defer close(a1Done)
	defer close(b1Done)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Statement A1")
		a1Done <- struct{}{}
		<-b1Done

		fmt.Println("Statement A2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Statement B1")
		b1Done <- struct{}{}
		<-a1Done

		fmt.Println("Statement B2")
	}()

	wg.Wait()
}
