package orchestrator

import "jarvis/internal/events"

func (o *Orchestrator) registerHandlers() {

	o.bus.Subscribe(
		events.WakewordDetected,
		o.handleWakewordDetected,
	)

	o.bus.Subscribe(
		events.RecordingFinished,
		o.handleRecordingFinished,
	)

	o.bus.Subscribe(
		events.TranscriptionCompleted,
		o.handleTranscriptionCompleted,
	)
}