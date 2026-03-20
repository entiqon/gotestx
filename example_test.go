package main

import (
	"os"

	internal "github.com/entiqon/gotestx/internal"
)

func Example_basic() {
	internal.Run([]string{"./..."}, os.Stdout, os.Stderr)
}

func Example_withCoverage() {
	internal.Run([]string{"-c"}, os.Stdout, os.Stderr)
}

func Example_withIgnore() {
	internal.Run([]string{"-I", "mock", "./..."}, os.Stdout, os.Stderr)
}

func Example_quiet() {
	internal.Run([]string{"-q"}, os.Stdout, os.Stderr)
}
