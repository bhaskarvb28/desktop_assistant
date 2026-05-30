package recording

import (
	"log"

	"time"

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
		events.StartRecording,
		m.handleStartRecording,
	)
}

func (m *Module) handleStartRecording(
	event events.Event,
) {

	log.Println(
		"[RECORDING] start requested",
	)

	session := &Session{
		ID: "session-1",

		State: SessionRecording,

		StartedAt: time.Now(),
	}

	go m.startSession(session)
}