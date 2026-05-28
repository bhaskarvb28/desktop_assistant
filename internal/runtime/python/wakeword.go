package python

import (
	"bufio"
	"fmt"
	"os/exec"
)

func StartWakeWordRuntime() error {

	cmd := exec.Command(
		"python",
		"runtimes/wakeword/main.py",
	)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	err = cmd.Start()

	if err != nil {
		return err
	}

	go func() {

		scanner := bufio.NewScanner(stdout)

		for scanner.Scan() {
			fmt.Println("[WAKEWORD]", scanner.Text())
		}
	}()

	return nil
}