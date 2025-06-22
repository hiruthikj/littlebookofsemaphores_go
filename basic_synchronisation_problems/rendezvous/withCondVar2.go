package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Added here to keep main clutter free
func init() {
	wg.Add(2)
}

// A1 -> B2
// A1 should be done before B2
// B1 -> A2
// B1 should be done before A
func main() {
	aReached, bReached := false, false
	rendev := sync.NewCond(&sync.Mutex{})

	go func() {
		defer wg.Done()
		fmt.Println("Statement A1")

		rendev.L.Lock()
		aReached = true
		fmt.Println("\tA reached rendev")

		for !bReached {
			fmt.Println("Waiting for B rendev")
			rendev.Wait()
		}
		fmt.Println("\tA wait done")
		rendev.L.Unlock()
		// Without signal, it will go to wait and all will be sleeping
		rendev.Signal()

		fmt.Println("Statement A2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Statement B1")

		rendev.L.Lock()
		bReached = true
		fmt.Println("\tB reached rendev")

		for !aReached {
			fmt.Println("\tWaiting for A rendev")
			rendev.Wait()
		}
		fmt.Println("\tB wait done")
		rendev.L.Unlock()
		rendev.Signal()

		fmt.Println("Statement B2")
	}()

	wg.Wait()
}
