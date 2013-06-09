package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"
)

var spin   string
var format string
var delay  int

func init() {
	flag.IntVar(&delay, "d", 100, "delay between `frames' in milliseconds")
	flag.StringVar(&spin, "s", "/-\\|", "animation to play")
	flag.StringVar(&format, "f", "%s",
		"format string for the animation. %s is replaced by each frame")

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

		fmt.Print(strings.Replace(format, "%s", string(spin[i]), -1))
		fmt.Print("\r")
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	cleanup()
}

func cleanup() {
	fmt.Println()
	os.Exit(0)
}
