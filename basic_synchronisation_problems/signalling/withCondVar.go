package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	signal := sync.NewCond(&sync.Mutex{})

	go func() {
		for range 30 {
			fmt.Println("Statement A1")
		}
		signal.L.Lock()
		defer signal.L.Unlock()

		signal.Signal()

	}()

	go func() {
		signal.L.Lock()
		defer signal.L.Unlock()

		signal.Wait()

		for range 30 {
			fmt.Println("Statement B1")
		}
	}()

	// using sleep for simplicity
	time.Sleep(1 * time.Second)
	os.Stdout.Sync()
}
