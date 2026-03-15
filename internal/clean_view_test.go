package gotestx_test

import (
	"bytes"
	"strings"
	"testing"

	gotestx "github.com/entiqon/gotestx/internal"
)

func TestFilterCleanViewOutput(t *testing.T) {
	input := `
ok   pkg/a
?    pkg/b [no test files]
ok   pkg/c
`

	var out bytes.Buffer

	gotestx.FilterCleanViewOutput(strings.NewReader(input), &out)

	result := out.String()

	if strings.Contains(result, "[no test files]") {
		t.Fatalf("expected '[no test files]' to be filtered")
	}

	if !strings.Contains(result, "pkg/a") || !strings.Contains(result, "pkg/c") {
		t.Fatalf("expected valid lines to remain")
	}
}
