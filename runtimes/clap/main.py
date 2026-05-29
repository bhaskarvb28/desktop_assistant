import pyaudio
import numpy as np
import time

last_clap = 0

CHUNK_SIZE = 1280
SAMPLE_RATE = 16000

audio = pyaudio.PyAudio()

stream = audio.open(
    format=pyaudio.paInt16,
    channels=1,
    rate=SAMPLE_RATE,
    input=True,
    frames_per_buffer=CHUNK_SIZE
)

while True:

    audio_chunk = stream.read(
        CHUNK_SIZE,
        exception_on_overflow=False
    )

    audio_data = np.frombuffer(
        audio_chunk,
        dtype=np.int16
    )

    volume = np.abs(audio_data).mean()

    now = time.time()

    if (
        volume > 1000 and
        now - last_clap > 0.5
    ):

        print("CLAP DETECTED")

        last_clap = now