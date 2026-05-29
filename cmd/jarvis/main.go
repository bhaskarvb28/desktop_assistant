package main

import (
	"log"

	"jarvis/internal/audio/recording"
	"jarvis/internal/events"
	"jarvis/internal/orchestrator"
	"jarvis/internal/runtime/python"
)

func main() {

	// --------------------------------------------------
	// Event Bus
	// --------------------------------------------------

	bus := events.NewBus()

	// --------------------------------------------------
	// Orchestrator
	// --------------------------------------------------

	orch := orchestrator.New(bus)
	orch.Start()

	// --------------------------------------------------
	// Modules
	// --------------------------------------------------

	recording.Register(bus)

	// --------------------------------------------------
	// Runtimes
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