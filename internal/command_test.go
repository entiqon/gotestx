package gotestx

import (
	"bytes"
	"runtime"
	"testing"
)

func TestCommandRunner(t *testing.T) {

	t.Run("Real", func(t *testing.T) {

		t.Run("Run", func(t *testing.T) {

			cmd := commandRunner("go", "version")

			var out bytes.Buffer
			var errBuf bytes.Buffer

			cmd.SetStdout(&out)
			cmd.SetStderr(&errBuf)

			if err := cmd.Run(); err != nil {
				t.Fatalf("expected command to run successfully: %v", err)
			}

			if out.Len() == 0 {
				t.Fatalf("expected stdout output")
			}
		})

		t.Run("Stdout", func(t *testing.T) {

			cmd := commandRunner("go", "version")

			var out bytes.Buffer
			cmd.SetStdout(&out)

			if err := cmd.Run(); err != nil {
				t.Fatalf("command failed: %v", err)
			}

			if out.Len() == 0 {
				t.Fatalf("expected stdout output")
			}
		})

		t.Run("Failure", func(t *testing.T) {

			var cmd Command

			if runtime.GOOS == "windows" {
				cmd = commandRunner("cmd", "/c", "exit", "1")
			} else {
				cmd = commandRunner("false")
			}

			err := cmd.Run()

			if err == nil {
				t.Fatalf("expected command to fail")
			}
		})
	})

	t.Run("CreatesCommand", func(t *testing.T) {

		cmd := commandRunner("go", "version")

		if cmd == nil {
			t.Fatalf("expected commandRunner to return command")
		}

		cmd.SetStdout(&bytes.Buffer{})
		cmd.SetStderr(&bytes.Buffer{})

		if err := cmd.Run(); err != nil {
			t.Fatalf("command should run: %v", err)
		}
	})
}
