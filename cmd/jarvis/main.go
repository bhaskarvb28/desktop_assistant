package main

import (
	"fmt"
	"log"

	"jarvis/internal/events"
	"jarvis/internal/runtime/python"
)

func main() {

	fmt.Println("STARTED")

	// --------------------------------------------------
	// Event Bus
	// --------------------------------------------------

	bus := events.NewBus()

	// --------------------------------------------------
	// Register Event Handlers
	// --------------------------------------------------

	registerWakewordHandlers(bus)

	// --------------------------------------------------
	// Start Runtimes
	// --------------------------------------------------

	err := python.StartWakeWordRuntime(bus)

	if err != nil {
		log.Fatal(err)
	}

	// --------------------------------------------------
	// Block Forever
	// --------------------------------------------------

	select {}
}

func registerWakewordHandlers(
	bus *events.Bus,
) {

	wakewordChannel := bus.Subscribe(
		events.WakewordDetected,
	)

	go func() {

		for event := range wakewordChannel {

			fmt.Println(
				"[EVENT] Wakeword detected:",
				event.Data,
			)
		}
	}()
}