package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gotest.tools/gotestsum/cmd"
)

const (
	flagComposeFileDefault = "docker-compose.yml"
	flagComposeFileUsage   = "set the docker compose file (defaults to docker-compose.yml)"

	flagCoverDefault = true
	flagCoverUsage   = "show code coverage percentage"

	flagRaceDefault = false
	flagRaceUsage   = "run race detection on tests"
)

func main() {
	var composeFile string
	flag.StringVar(&composeFile, "compose_file", flagComposeFileDefault, flagComposeFileUsage)
	flag.StringVar(&composeFile, "cf", flagComposeFileDefault, flagComposeFileUsage)

	var rawCmd string
	flag.StringVar(&rawCmd, "raw", "", "input a custom command, this will override any other test command arguments (ex: -cover, -race)")

	var cover bool
	flag.BoolVar(&cover, "cover", flagCoverDefault, flagCoverUsage)
	flag.BoolVar(&cover, "c", flagCoverDefault, flagCoverUsage)

	var race bool
	flag.BoolVar(&race, "race", flagRaceDefault, flagRaceUsage)
	flag.BoolVar(&race, "r", flagRaceDefault, flagRaceUsage)

	flag.Parse()

	cmdArgs := []string{"--"}
	if rawCmd != "" {
		cmdArgs = append(cmdArgs, strings.Split(rawCmd, " ")...)
	} else {
		if cover {
			cmdArgs = append(cmdArgs, "-cover")
		}

		if race {
			cmdArgs = append(cmdArgs, "-race")
		}

		if len(flag.Args()) == 0 {
			cmdArgs = append(cmdArgs, "./...")
		} else {
			cmdArgs = append(cmdArgs, flag.Args()...)
		}
	}

	if err := setupDocker(composeFile); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error setting up docker compose: %v", err)
		os.Exit(1)
	}

	if err := cmd.Run("go_test", cmdArgs); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error running test command: %v", err)
		os.Exit(1)
	}
}

func setupDocker(composeFile string) error {
	_, err := os.Stat(composeFile)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("no compose file found, skipping docker image setup...")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check for existence of docker compose file: %v", err)
	}

	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d", "--wait")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
