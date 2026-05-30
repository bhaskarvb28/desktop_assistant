package recording

import "time"

type Session struct {
	ID string

	State SessionState

	StartedAt time.Time

	OutputFile string
}