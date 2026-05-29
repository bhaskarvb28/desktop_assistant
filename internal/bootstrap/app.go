package bootstrap

import (
	"jarvis/internal/audio/recording"
	"jarvis/internal/events"
	"jarvis/internal/orchestrator"
	"jarvis/internal/runtime/python"
	"jarvis/internal/state"
)

type App struct {
	bus *events.Bus
	orch *orchestrator.Orchestrator
	state *state.Manager
}

func New() (*App, error) {

	bus := events.NewBus()

	stateManager := state.NewManager()

	orch := orchestrator.New(
		bus,
		stateManager,
	)

	return &App{
		bus:  bus,
		orch: orch,
		state: stateManager,
	}, nil
}

func (a *App) Start() error {

	// orchestrator
	a.orch.Start()

	// modules
	// When this scales need separate modules.go
	recording.Register(a.bus)

	// runtimes
	// When this scales need separate modules.go
	err := python.StartWakeWordRuntime(a.bus)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) Wait() {
	select {}
}