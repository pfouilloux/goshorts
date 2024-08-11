package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
TODO: This is a work in progress. The goal is to create a tool that will list dependencies, list dependencies that are out of date, and update dependencies.
Commands:
- list: list dependencies
	- all: list all dependencies
	- direct: list direct dependencies
	- indirect: list indirect dependencies
	- outdated: list outdated dependencies
	- help: show help
- update: update dependencies
	- all: update all dependencies
	- direct: update direct dependencies
	- indirect: update indirect dependencies
	- outdated: update outdated dependencies
	- help: show help
- help: show help
*/

func main() {
	// go list -m -f '{{if not (or .Indirect .Main)}}{{.Path}}{{end}}' all
	var packages strings.Builder

	cmd := exec.Command("go", "list", "-m", "-f", "{{if not (or .Indirect .Main)}}{{.Path}}{{end}}", "all")
	cmd.Stdout = &packages
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error running go list command: %v", err)
		os.Exit(1)
	}

	strings.Split(packages.String(), "\n")

}
