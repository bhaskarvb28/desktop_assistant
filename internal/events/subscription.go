package events

type Subscription struct {
	ID        uint64
	EventType EventType
	Handler   Handler
}