package main

import (
	"log"

	"jarvis/internal/runtime/python"
)

func main() {
	err := python.StartWakeWordRuntime()

	if err != nil {
		log.Fatal(err)
	}

	select {}
}