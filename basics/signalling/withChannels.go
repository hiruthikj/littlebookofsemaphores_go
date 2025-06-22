package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	signal := make(chan struct{})

	go func() {
		for range 30 {
			fmt.Println("Statement A1")
		}
		signal <- struct{}{}
	}()

	go func() {
		<-signal
		for range 30 {
			fmt.Println("Statement B1")
		}
	}()

	// Ideally should never use as it will mask deadlocks on incorrect code
	time.Sleep(1 * time.Second)
	os.Stdout.Sync()
}
