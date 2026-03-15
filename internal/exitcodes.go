package gotestx

// Exit codes used internally by GoTestX CLI execution flow.
const (
	// ExitContinue indicates that execution should continue normally.
	// This is an internal sentinel value used by ResolveOptions to signal
	// that no early exit (help/version/error) was requested.
	ExitContinue = -1

	// ExitOK indicates the command completed successfully without
	// executing the test workflow (e.g. help or version output).
	ExitOK = 0

	// ExitUsage indicates a CLI usage error such as an unknown flag.
	ExitUsage = 2
)
