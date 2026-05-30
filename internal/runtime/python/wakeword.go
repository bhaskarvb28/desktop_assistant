package python

import (
	"bufio"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"jarvis/internal/events"
	"jarvis/internal/config"
)

func StartWakeWordRuntime(
	bus *events.Bus,
) error {

	root := config.JarvisRoot()

	pythonPath := filepath.Join(
		root,
		"runtimes",
		"wakeword",
		"venv",
		"Scripts",
		"python.exe",
	)

	scriptPath := filepath.Join(
		root,
		"runtimes",
		"wakeword",
		"main.py",
	)

	log.Println("python:", pythonPath)
	log.Println("script:", scriptPath)

	cmd := exec.Command(
		pythonPath,
		"-u",
		scriptPath,
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
					Payload: wakeword,
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