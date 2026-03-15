package gotestx

import "runtime"

// getGOOS returns the current operating system.
//
// It is defined as a variable (instead of calling runtime.GOOS directly)
// so tests can override it to simulate different environments.
var getGOOS = func() string {
	return runtime.GOOS
}
