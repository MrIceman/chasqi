package main

import (
	"chasqi/engine"
	"chasqi/processor"
)

func main() {
	tree := processor.GetNavigationTree("./processor/fixture/navigation.yaml")

	scheduler := engine.NewScheduler(tree)
	debugChan := make(chan string)

	scheduler.Schedule(
		debugChan,
	)

	println("starting ")

	scheduler.Start()

	for {
		select {
		case debug := <-debugChan:
			println("Received debug message:" + debug)
		}
	}
}
