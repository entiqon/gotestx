package gotestx

import (
	"fmt"
	"io"
	"runtime"
)

// versionInfo prints detailed version information for GoTestX.
//
// The output includes the tool name, description, author, version,
// Git commit, build date, and the current processor architecture
// and operating system.
func versionInfo(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"%s\n\n%s\n%s\nAuthor: %s\nVersion: %s\nCommit: %s\nBuilt: %s\nProcessor: %s (%s)\n",
		ToolName,
		Description,
		ToolName,
		Author,
		Version,
		GitCommit,
		BuildDate,
		runtime.GOARCH,
		runtime.GOOS,
	)
}
