import openwakeword
from openwakeword.model import Model

import pyaudio
import numpy as np
import sys
import time

CHUNK_SIZE = 1280
SAMPLE_RATE = 16000

openwakeword.utils.download_models()

model = Model(
    inference_framework="onnx"
)

audio = pyaudio.PyAudio()

stream = audio.open(
    format=pyaudio.paInt16,
    channels=1,
    rate=SAMPLE_RATE,
    input=True,
    frames_per_buffer=CHUNK_SIZE
)

print("wakeword runtime started")
print("listening...")
sys.stdout.flush()

last_detection_time = 0
cooldown_seconds = 2

while True:

    audio_chunk = stream.read(
        CHUNK_SIZE,
        exception_on_overflow=False
    )

    audio_data = np.frombuffer(
        audio_chunk,
        dtype=np.int16
    )

    prediction = model.predict(audio_data)

    for wakeword, score in prediction.items():

        current_time = time.time()

        if (
            score > 0.7 and
            current_time - last_detection_time > cooldown_seconds
        ):

            print(f"detected:{wakeword}")
            sys.stdout.flush()

            last_detection_time = current_time