package gotestx

import "testing"

func TestFilterPackages(t *testing.T) {
	tests := []struct {
		name     string
		pkgs     []string
		patterns []string
		expect   []string
	}{
		{
			name:     "NoPatternsReturnsAll",
			pkgs:     []string{"pkg/a", "pkg/b"},
			patterns: nil,
			expect:   []string{"pkg/a", "pkg/b"},
		},
		{
			name:     "EmptyPackages",
			pkgs:     []string{},
			patterns: []string{"mock"},
			expect:   []string{},
		},
		{
			name:     "FilterSingleMatch",
			pkgs:     []string{"pkg/a", "pkg/mock", "pkg/b"},
			patterns: []string{"mock"},
			expect:   []string{"pkg/a", "pkg/b"},
		},
		{
			name:     "FilterMultipleMatches",
			pkgs:     []string{"pkg/a/mock", "pkg/b/testkit", "pkg/c"},
			patterns: []string{"mock", "testkit"},
			expect:   []string{"pkg/c"},
		},
		{
			name:     "FilterExactPath",
			pkgs:     []string{"pkg/outport/testkit", "pkg/a"},
			patterns: []string{"outport/testkit"},
			expect:   []string{"pkg/a"},
		},
		{
			name:     "NoMatchReturnsAll",
			pkgs:     []string{"pkg/a", "pkg/b"},
			patterns: []string{"mock"},
			expect:   []string{"pkg/a", "pkg/b"},
		},
		{
			name:     "AllFiltered",
			pkgs:     []string{"pkg/mock", "pkg/testkit"},
			patterns: []string{"mock", "testkit"},
			expect:   []string{},
		},
		{
			name:     "OrderPreserved",
			pkgs:     []string{"pkg/a", "pkg/mock", "pkg/b", "pkg/testkit"},
			patterns: []string{"mock", "testkit"},
			expect:   []string{"pkg/a", "pkg/b"},
		},
		{
			name:     "PatternWithSpaces",
			pkgs:     []string{"pkg/mock"},
			patterns: []string{"  mock  "},
			expect:   []string{},
		},
		{
			name:     "EmptyPatternEntriesIgnored",
			pkgs:     []string{"pkg/mock", "pkg/a"},
			patterns: []string{"", "   ", "mock"},
			expect:   []string{"pkg/a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterPackages(tt.pkgs, tt.patterns)

			if len(got) != len(tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, got)
			}

			for i := range got {
				if got[i] != tt.expect[i] {
					t.Fatalf("expected %v, got %v", tt.expect, got)
				}
			}
		})
	}
}
