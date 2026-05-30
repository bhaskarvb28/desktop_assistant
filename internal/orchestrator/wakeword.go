package orchestrator

import (
	"log"

	"jarvis/internal/events"
	"jarvis/internal/state"
)

func (o *Orchestrator) handleWakewordDetected(
	event events.Event,
) {

	log.Println(
		"[ORCHESTRATOR] wakeword detected:",
		event.Payload,
	)

	// --------------------------------------------------
	// Ignore if busy
	// --------------------------------------------------

	if o.state.Get() != state.Idle {

		log.Println(
			"[ORCHESTRATOR] assistant busy",
		)

		return
	}

	// --------------------------------------------------
	// Transition State
	// --------------------------------------------------

	o.state.Set(
		state.Listening,
	)

	// --------------------------------------------------
	// Trigger Recording
	// --------------------------------------------------

	o.bus.Publish(events.Event{
		Type: events.StartRecording,
	})
}