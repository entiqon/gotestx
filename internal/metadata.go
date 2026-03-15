package gotestx

// Tool metadata used across CLI output, diagnostics, and build reporting.
//
// Version, GitCommit, and BuildDate may be overridden at build time using
// Go linker flags (-ldflags). When not provided, they fall back to the
// default development values defined below.
var (
	// Version identifies the GoTestX release version.
	//
	// By default it is set to "dev" for local builds. Release pipelines
	// typically override this value using:
	//
	//   go build -ldflags "-X github.com/entiqon/gotestx/internal.Version=vX.Y.Z"
	Version = "dev"

	// GitCommit identifies the source control revision used to build
	// the binary. It is normally injected during CI builds.
	GitCommit = "none"

	// BuildDate records the UTC timestamp when the binary was built.
	// This value is also commonly injected by CI pipelines.
	BuildDate = "unknown"

	// ToolName is the CLI command name used when printing help
	// messages or diagnostics.
	ToolName = "GoTestX"

	// Author identifies the project maintainers or organization.
	Author = "Entiqon Labs Team"

	// Description provides a short explanation of the tool purpose
	// displayed in CLI help output.
	Description = "Go Test eXtended tool with coverage support"
)
