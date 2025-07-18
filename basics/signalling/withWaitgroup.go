package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 30; i++ {
			fmt.Println("Statement A1")
		}
	}()

	go func() {
		wg.Wait()
		for i := 0; i < 30; i++ {
			fmt.Println("Statement B1")
		}
	}()

	// Ideally should never use as it will mask deadlocks on incorrect code
	time.Sleep(1 * time.Second)
	os.Stdout.Sync()
}
