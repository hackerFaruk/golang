package main

import (
	"fmt"
	"time"
)

func waitMe() {
	fmt.Println("This is before the delay.")

	// Delay for 2 seconds
	duration := 2 * time.Second
	time.Sleep(duration)

	fmt.Println("This is after the delay.")
}
