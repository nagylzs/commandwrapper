package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var Command string = ""

func main() {
	if Command == "" {
		os.Stderr.WriteString("bad wrapper: command was not set at compile time\n")
		os.Exit(1)
	}

	cmd := exec.Command(Command, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		os.Stderr.WriteString("Failed to start: " + err.Error() + "\n")
		os.Exit(1)
	}

	// Forward all signals to the child process
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for sig := range sigs {
			_ = cmd.Process.Signal(sig)
		}
	}()

	// Wait for the command to finish
	err := cmd.Wait()

	// Exit with the same code
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
	}

	if err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Exit(0)
}
