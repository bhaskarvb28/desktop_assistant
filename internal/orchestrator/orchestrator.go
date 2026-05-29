package orchestrator

import (
	"jarvis/internal/events"
)

type Orchestrator struct {
	bus *events.Bus
}

func New(
	bus *events.Bus,
) *Orchestrator {

	return &Orchestrator{
		bus: bus,
	}
}

func (o *Orchestrator) Start() {

	o.registerWakewordHandlers()
}