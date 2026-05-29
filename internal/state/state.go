package state

type AssistantState string

const (
	Idle       AssistantState = "IDLE"
	Listening  AssistantState = "LISTENING"
	Processing AssistantState = "PROCESSING"
	Responding AssistantState = "RESPONDING"
)