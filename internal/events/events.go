package events

type EventType string

const (
	WakewordDetected EventType = "WAKEWORD_DETECTED"
	StartRecording	 EventType = "START_RECORDING"
)

type Event struct {
	Type EventType
	Data any
}