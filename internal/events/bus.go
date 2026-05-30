package events

import (
	"sync"
	"sync/atomic"
)

type Bus struct {
	mu sync.RWMutex

	nextID uint64

	subscribers map[EventType]map[uint64]Handler
}

func NewBus() *Bus {

	return &Bus{
		subscribers: make(
			map[EventType]map[uint64]Handler,
		),
	}
}

func (b *Bus) Subscribe(
	eventType EventType,
	handler Handler,
) *Subscription {

	b.mu.Lock()
	defer b.mu.Unlock()

	id := atomic.AddUint64(
		&b.nextID,
		1,
	)

	if b.subscribers[eventType] == nil {

		b.subscribers[eventType] = make(
			map[uint64]Handler,
		)
	}

	b.subscribers[eventType][id] = handler

	return &Subscription{
		ID:        id,
		EventType: eventType,
		Handler:   handler,
	}
}

func (b *Bus) Unsubscribe(
	sub *Subscription,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	handlers, exists := b.subscribers[sub.EventType]

	if !exists {
		return
	}

	delete(
		handlers,
		sub.ID,
	)

	if len(handlers) == 0 {

		delete(
			b.subscribers,
			sub.EventType,
		)
	}
}

func (b *Bus) Publish(
	event Event,
) {

	b.mu.RLock()

	handlersMap := b.subscribers[event.Type]

	handlers := make(
		[]Handler,
		0,
		len(handlersMap),
	)

	for _, handler := range handlersMap {
		handlers = append(
			handlers,
			handler,
		)
	}

	b.mu.RUnlock()

	for _, handler := range handlers {

		go handler(event)
	}
}