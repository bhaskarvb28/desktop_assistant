package recording

import (
	"log"

	"jarvis/internal/events"
)

func Register(
	bus *events.Bus,
) {

	recordingChannel := bus.Subscribe(
		events.StartRecording,
	)

	go func() {

		for range recordingChannel {

			log.Println(
				"[RECORDING] recording started",
			)
		}
	}()
}