package main

import (
	"log"

	"jarvis/internal/bootstrap"
)

func main() {
	
	app, err := bootstrap.New()

	if err != nil {
		log.Fatal(err)
	}

	err = app.Start()

	if err != nil {
		log.Fatal(err)
	}

	app.Wait()
}