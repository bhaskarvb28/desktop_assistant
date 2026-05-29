package events

import "sync"

type Bus struct {
	mu sync.RWMutex

	subscribers map[EventType][]chan Event
}

func NewBus() *Bus {

	return &Bus{
		subscribers: make(
			map[EventType][]chan Event,
		),
	}
}

func (b *Bus) Subscribe(
	eventType EventType,
) <-chan Event {

	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Event, 10)

	b.subscribers[eventType] = append(
		b.subscribers[eventType],
		ch,
	)

	return ch
}

func (b *Bus) Publish(event Event) {

	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subscribers[event.Type] {

		select {

		case ch <- event:

		default:
			// subscriber is slow
			// skip event
		}
	}
}