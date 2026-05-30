package orchestrator

import (
	"log"

	"jarvis/internal/events"
	"jarvis/internal/state"
)

func (o *Orchestrator) handleTranscriptionCompleted(
	event events.Event,
) {

	log.Println(
		"[ORCHESTRATOR] user said:",
		event.Payload,
	)

	o.state.Set(
		state.Idle,
	)
}