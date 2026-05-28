from pathlib import Path

folders = [
    "cmd/jarvis",

    "internal/app",
    "internal/orchestrator",
    "internal/events",
    "internal/config",
    "internal/logging",
    "internal/shared",

    "internal/audio/wakeword",
    "internal/audio/recording",
    "internal/audio/stt",
    "internal/audio/tts",

    "internal/ai/ollama",
    "internal/ai/prompts",
    "internal/ai/tools",

    "internal/runtime/python",
    "internal/runtime/manager",

    "internal/system",

    "runtimes/wakeword",
    "runtimes/whisper",
    "runtimes/vision",

    "models",
    "scripts",
    "build",
    "tmp",
]

for folder in folders:
    Path(folder).mkdir(parents=True, exist_ok=True)

print("Jarvis folder structure created.")