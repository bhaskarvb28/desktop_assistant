package recording

import (
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func writeWavFile(
	filename string,
	samples []int,
	sampleRate int,
) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := wav.NewEncoder(
		file,
		sampleRate,
		16,
		1,
		1,
	)

	buffer := &audio.IntBuffer{
		Data: samples,

		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  sampleRate,
		},
	}

	err = encoder.Write(buffer)

	if err != nil {
		return err
	}

	return encoder.Close()
}