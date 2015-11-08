package main

import (
	"fmt"
	"time"
	"os/exec"
)

// The Uptime function just runs uptime and returns the output
func Uptime() string {
	out, err := exec.Command("uptime").Output()
	if err != nil {
		return "Unable to run uptime\n"
	}
	return string(out)
}

// Snoozer's  sends the current time to another goroutine every second.
// The presence of data on that channel is intended to indicate to that
// other goroutine that it's time to do work.
func Snoozer(...) {
	for {
		time.Sleep(time.Second)
		// Attempt to send time.Now to the c chan, but don't wait for it.
		// If we wait for it, we might be stuck for an indefinite amount of
		// time and fail at our goal of sending it every second.
		...
		...
		...
		}
	}
}

func main() {
	// Create our channel for Snoozer to user
	t := ...
	// Start up Snoozer
	go Snoozer(t)
	for {
		// Get data from Snoozer, or wait if Snoozer hasn't said it's time yet
		ct := ...
		// Output current time, and uptime output
		fmt.Printf("%s : %s", ct, Uptime())
	}
}
