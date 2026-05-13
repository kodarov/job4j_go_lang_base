package main

import (
	"job4j.ru/go-lang-base/internal/tracker"
)

func main() {
	ui := tracker.UI{
		In:      tracker.ConsoleInput{},
		Out:     tracker.ConsoleOutput{},
		Tracker: tracker.NewTracker(),
	}
	ui.Run()
}
