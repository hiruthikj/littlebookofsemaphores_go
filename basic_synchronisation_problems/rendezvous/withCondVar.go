package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func init() {
	wg.Add(2)
}

// A1 -> B2
// A1 should be done before B2
// B1 -> A2
// B1 should be done before A
func main() {
	bReached := false
	rendev := sync.NewCond(&sync.Mutex{})

	go func() {
		defer wg.Done()
		fmt.Println("Statement A1")
		// aReached = true

		rendev.L.Lock()
		for !bReached {
			rendev.Wait()
		}
		rendev.L.Unlock()
		// Making other to wait without any check
		rendev.Signal()

		fmt.Println("Statement A2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Statement B1")

		// If we don't surround with lock and unlock - we get fatal "unlock of unlocked mutex"
		rendev.L.Lock()
		// Updating this should be with a lock
		bReached = true
		rendev.Wait()
		rendev.L.Unlock()

		fmt.Println("Statement B2")
	}()

	wg.Wait()
}
