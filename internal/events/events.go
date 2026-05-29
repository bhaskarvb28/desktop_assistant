package events

type EventType string

const (
	WakewordDetected EventType = "WAKEWORD_DETECTED"
)

type Event struct {
	Type EventType
	Data any
}