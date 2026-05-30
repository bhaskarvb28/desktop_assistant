package orchestrator

import (
	"log"

	"jarvis/internal/events"
	"jarvis/internal/state"
)

func (o *Orchestrator) handleRecordingFinished(
	event events.Event,
) {

	log.Println(
		"[ORCHESTRATOR] recording finished",
	)

	o.state.Set(
		state.Processing,
	)
}