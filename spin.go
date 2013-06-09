package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var spin = []string{ "/", "-", "\\", "|" }

var delay int // Delay in milliseconds

func init() {
	flag.IntVar(&delay, "d", 100, "delay between `frames' in milliseconds")

	flag.Parse()
}

func main() {
	// Make sure a newline is printed even on SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		cleanup()
	}()

	for i := 0;; i++ {
		if i == len(spin) {
			i = 0
		}

		fmt.Print(spin[i])
		fmt.Print("\r")
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	cleanup()
}

func cleanup() {
	fmt.Println()
	os.Exit(0)
}
