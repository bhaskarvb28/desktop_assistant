package stt

import (
	"jarvis/internal/events"
)

type Module struct {
	bus *events.Bus
}

func New(
	bus *events.Bus,
) *Module {

	return &Module{
		bus: bus,
	}
}

func (m *Module) Start() {

	m.registerHandlers()
}

func (m *Module) registerHandlers() {

	m.bus.Subscribe(
		events.RecordingFinished,
		m.handleRecordingFinished,
	)
}

func (m *Module) handleRecordingFinished(
	event events.Event,
) {

	audioFile := event.Payload.(string)

	go m.transcribe(audioFile)
}