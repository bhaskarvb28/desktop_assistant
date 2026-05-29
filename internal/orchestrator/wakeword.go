package orchestrator

import (
	"log"

	"jarvis/internal/events"
)

func (o *Orchestrator) registerWakewordHandlers() {

	wakewordChannel := o.bus.Subscribe(
		events.WakewordDetected,
	)

	go func() {

		for event := range wakewordChannel {

			log.Println(
				"[ORCHESTRATOR] wakeword detected:",
				event.Data,
			)

			// --------------------------------------------------
			// Trigger Recording Workflow
			// --------------------------------------------------

			o.bus.Publish(events.Event{
				Type: events.StartRecording,
				Data: nil,
			})
		}
	}()
}