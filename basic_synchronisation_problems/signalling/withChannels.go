package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	signal := make(chan struct{})

	go func() {
		for _ = range 30 {
			fmt.Println("Statement A1")
		}
		signal <- struct{}{}
	}()

	go func() {
		<-signal
		for _ = range 30 {
			fmt.Println("Statement B1")
		}
	}()

	time.Sleep(1 * time.Second)
	os.Stdout.Sync()

}
