package gotestx

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

type fakeCommand struct {
	runErr error
	output string
	stdout io.Writer
	stderr io.Writer
}

func (f *fakeCommand) Run() error {

	if f.stdout != nil {
		f.stdout.Write([]byte(f.output))
	}

	return f.runErr
}

func (f *fakeCommand) SetStdout(w io.Writer) {
	f.stdout = w
}

func (f *fakeCommand) SetStderr(w io.Writer) {
	f.stderr = w
}

func run(t *testing.T, args ...string) (string, string, int) {

	t.Helper()

	var out bytes.Buffer
	var err bytes.Buffer

	code := Run(args, &out, &err)

	return out.String(), err.String(), code
}

func TestHelp(t *testing.T) {

	out, _, code := run(t, "-h")

	if code != 0 {
		t.Fatalf("expected exit 0")
	}

	if !strings.Contains(out, "Usage") {
		t.Fatalf("expected usage")
	}
}

func TestVersion(t *testing.T) {

	out, _, code := run(t, "-v")

	if code != 0 {
		t.Fatalf("expected exit 0")
	}

	if !strings.Contains(out, Version) {
		t.Fatalf("expected version output")
	}
}

func TestGoTestFailure(t *testing.T) {

	old := commandRunner
	defer func() { commandRunner = old }()

	commandRunner = func(name string, args ...string) Command {

		return &fakeCommand{
			runErr: errors.New("fail"),
		}
	}

	_, err, code := run(t, "./")

	if code == 0 {
		t.Fatalf("expected failure")
	}

	if !strings.Contains(err, "go test failed") {
		t.Fatalf("expected go test failed error")
	}
}

func TestRun(t *testing.T) {
	t.Run("Packages", func(t *testing.T) {
		t.Run("Invalid", func(t *testing.T) {
			var out bytes.Buffer
			var err bytes.Buffer

			code := Run([]string{"./does-not-exist"}, &out, &err)

			if code == ExitOK {
				t.Fatalf("expected failure")
			}
		})

		t.Run("Success", func(t *testing.T) {
			orig := commandRunner
			defer func() { commandRunner = orig }()

			commandRunner = func(name string, args ...string) Command {
				return &fakeCommand{}
			}

			var out bytes.Buffer

			code := Run([]string{"./..."}, &out, io.Discard)

			if code != ExitOK {
				t.Fatalf("expected success")
			}
		})
	})

	t.Run("Coverage", func(t *testing.T) {
		orig := commandRunner
		defer func() { commandRunner = orig }()

		commandRunner = func(name string, args ...string) Command {
			return &fakeCommand{}
		}

		var out bytes.Buffer

		code := Run([]string{"-c"}, &out, io.Discard)

		if code != ExitOK {
			t.Fatalf("expected success")
		}
	})
}
