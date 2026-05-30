package stt

import (
	"log"
	"os/exec"
	"strings"

	"jarvis/internal/events"
)

func (m *Module) transcribe(
	audioFile string,
) {

	log.Println(
		"[STT] transcribing:",
		audioFile,
	)

	cmd := exec.Command(
		"runtimes/whisper/build/bin/whisper-cli.exe",

		"-m",
		"runtimes/whisper/models/ggml-base.en.bin",

		"-f",
		audioFile,

		"-nt",
		"-np",
	)

	output, err := cmd.Output()

	if err != nil {

		log.Println(
			"[STT] transcription failed:",
			err,
		)

		m.bus.Publish(events.Event{
			Type: events.TranscriptionFailed,

			Payload: err.Error(),
		})

		return
	}

	text := strings.TrimSpace(
		string(output),
	)

	log.Println(
		"[STT] transcript:",
		text,
	)

	m.bus.Publish(events.Event{
		Type: events.TranscriptionCompleted,

		Payload: text,
	})
}