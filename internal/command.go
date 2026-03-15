package gotestx

import (
	"io"
	"os/exec"
)

// Command defines the minimal interface required to execute
// an external command. It allows command execution to be
// replaced during tests.
type Command interface {
	Run() error
	SetStdout(io.Writer)
	SetStderr(io.Writer)
}

// realCommand wraps exec.Cmd to implement Command.
type realCommand struct {
	cmd *exec.Cmd
}

func (c *realCommand) Run() error {
	return c.cmd.Run()
}

func (c *realCommand) SetStdout(w io.Writer) {
	c.cmd.Stdout = w
}

func (c *realCommand) SetStderr(w io.Writer) {
	c.cmd.Stderr = w
}

// commandRunner creates commands used by the runner.
//
// It defaults to exec.Command but can be overridden in tests
// to simulate command execution.
var commandRunner = func(name string, args ...string) Command {
	return &realCommand{
		cmd: exec.Command(name, args...),
	}
}
