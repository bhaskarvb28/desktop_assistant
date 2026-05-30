package events

type EventType string

const (

	// --------------------------------------------------
	// Wakeword
	// --------------------------------------------------

	WakewordDetected EventType = "WAKEWORD_DETECTED"

	// --------------------------------------------------
	// Recording
	// --------------------------------------------------

	StartRecording EventType = "START_RECORDING"

	RecordingStarted EventType = "RECORDING_STARTED"

	RecordingFinished EventType = "RECORDING_FINISHED"

	RecordingFailed EventType = "RECORDING_FAILED"

	// --------------------------------------------------
	// STT
	// --------------------------------------------------

	TranscriptionCompleted EventType = "TRANSCRIPTION_COMPLETED"

	TranscriptionFailed EventType = "TRANSCRIPTION_FAILED"
)