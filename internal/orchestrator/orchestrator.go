package orchestrator

import (
	"jarvis/internal/events"
	"jarvis/internal/state"
)

type Orchestrator struct {
	bus *events.Bus

	state *state.Manager
}

func New(
	bus *events.Bus,
	stateManager *state.Manager,
) *Orchestrator {

	return &Orchestrator{
		bus: bus,
		state: stateManager,
	}
}

func (o *Orchestrator) Start() {

	o.registerWakewordHandlers()
}