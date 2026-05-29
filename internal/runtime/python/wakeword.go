package python

import (
	"bufio"
	"log"
	"os/exec"
	"strings"

	"jarvis/internal/events"
)

func StartWakeWordRuntime(
	bus *events.Bus,
) error {

	cmd := exec.Command(
		"runtimes/wakeword/venv/Scripts/python.exe",
		"-u",
		"runtimes/wakeword/main.py",
	)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()

	if err != nil {
		return err
	}

	err = cmd.Start()

	if err != nil {
		return err
	}

	// --------------------------------------------------
	// stdout reader
	// --------------------------------------------------

	go func() {

		scanner := bufio.NewScanner(stdout)

		scanner.Buffer(
			make([]byte, 1024),
			1024*1024,
		)

		for scanner.Scan() {

			line := scanner.Text()

			log.Println(
				"[WAKEWORD]",
				line,
			)

			if strings.HasPrefix(
				line,
				"detected:",
			) {

				wakeword := strings.TrimPrefix(
					line,
					"detected:",
				)

				bus.Publish(events.Event{
					Type: events.WakewordDetected,
					Data: wakeword,
				})
			}
		}

		if err := scanner.Err(); err != nil {

			log.Println(
				"wakeword stdout scanner error:",
				err,
			)
		}
	}()

	// --------------------------------------------------
	// stderr reader
	// --------------------------------------------------

	go func() {

		scanner := bufio.NewScanner(stderr)

		scanner.Buffer(
			make([]byte, 1024),
			1024*1024,
		)

		for scanner.Scan() {

			log.Println(
				"[WAKEWORD ERROR]",
				scanner.Text(),
			)
		}

		if err := scanner.Err(); err != nil {

			log.Println(
				"wakeword stderr scanner error:",
				err,
			)
		}
	}()

	// --------------------------------------------------
	// process watcher
	// --------------------------------------------------

	go func() {

		err := cmd.Wait()

		if err != nil {

			log.Println(
				"wakeword runtime crashed:",
				err,
			)

			return
		}

		log.Println(
			"wakeword runtime exited",
		)
	}()

	return nil
}