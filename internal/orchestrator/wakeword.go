package orchestrator

import (
	"log"

	"jarvis/internal/events"
	"jarvis/internal/state"
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
			// Ignore if assistant busy
			// --------------------------------------------------

			if o.state.Get() != state.Idle {

				log.Println(
					"[ORCHESTRATOR] assistant busy",
				)

				continue
			}

			// --------------------------------------------------
			// Transition State
			// --------------------------------------------------

			o.state.Set(
				state.Listening,
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