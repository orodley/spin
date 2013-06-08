package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
)

var spin = []string{ "/", "-", "\\", "|" }

func main() {
	// Make sure a newline is printed even on SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<- c
		cleanup()
	}()

	for i := 0;; i++ {
		if i == len(spin) {
			i = 0
		}

		fmt.Print(spin[i])
		fmt.Print("\r")
		time.Sleep(100 * time.Millisecond)
	}

	cleanup()
}

func cleanup() {
	fmt.Println()
	os.Exit(0)
}
