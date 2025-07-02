package main

import (
	"bytes"
	"fmt"
	"sync"
)

func printBytes(wg *sync.WaitGroup, byteSlice []byte) {
	go func() {
		defer wg.Done()
		var buff = bytes.Buffer{}

		for _, b := range byteSlice {
			fmt.Fprintf(&buff, "%c", b)
		}

		fmt.Println(buff.String())

	}()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	str := []byte("golang")

	// Slices passed is confined
	printBytes(&wg, str[:3])
	printBytes(&wg, str[3:])

	wg.Wait()
}
