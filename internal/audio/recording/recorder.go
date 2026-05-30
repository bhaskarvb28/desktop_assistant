package recording

import (
	"fmt"
	"log"
	"time"

	"jarvis/internal/events"

	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate = 16000
	duration   = 5
)

func (m *Module) startSession(
	session *Session,
) {

	log.Println(
		"[RECORDING] recording...",
	)

	// --------------------------------------------------
	// Initialize PortAudio
	// --------------------------------------------------

	err := portaudio.Initialize()

	if err != nil {

		log.Println(
			"portaudio init failed:",
			err,
		)

		return
	}

	defer portaudio.Terminate()

	// --------------------------------------------------
	// Audio Buffer
	// --------------------------------------------------

	buffer := make(
		[]int16,
		64,
	)

	var recordedSamples []int

	// --------------------------------------------------
	// Open Input Stream
	// --------------------------------------------------

	stream, err := portaudio.OpenDefaultStream(
		1,
		0,
		float64(sampleRate),
		len(buffer),
		&buffer,
	)

	if err != nil {

		log.Println(
			"stream open failed:",
			err,
		)

		return
	}

	defer stream.Close()

	// --------------------------------------------------
	// Start Stream
	// --------------------------------------------------

	err = stream.Start()

	if err != nil {

		log.Println(
			"stream start failed:",
			err,
		)

		return
	}

	defer stream.Stop()

	// --------------------------------------------------
	// Record Audio
	// --------------------------------------------------

	start := time.Now()

	for {

		if time.Since(start) >=
			duration*time.Second {

			break
		}

		err := stream.Read()

		if err != nil {

			log.Println(
				"stream read failed:",
				err,
			)

			return
		}

		for _, sample := range buffer {

			recordedSamples = append(
				recordedSamples,
				int(sample),
			)
		}
	}

	// --------------------------------------------------
	// Save WAV File
	// --------------------------------------------------

	filename := fmt.Sprintf(
		"tmp/recording_%d.wav",
		time.Now().Unix(),
	)

	err = writeWavFile(
		filename,
		recordedSamples,
		sampleRate,
	)

	if err != nil {

		log.Println(
			"wav write failed:",
			err,
		)

		return
	}

	log.Println(
		"[RECORDING] saved:",
		filename,
	)

	session.State = SessionFinished

	// --------------------------------------------------
	// Publish Event
	// --------------------------------------------------

	m.bus.Publish(events.Event{
		Type: events.RecordingFinished,

		Payload: filename,
	})
}