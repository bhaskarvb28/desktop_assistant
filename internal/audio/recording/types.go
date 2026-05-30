package recording

type SessionState string

const (
	SessionIdle SessionState = "IDLE"

	SessionRecording SessionState =
		"RECORDING"

	SessionFinished SessionState =
		"FINISHED"
)