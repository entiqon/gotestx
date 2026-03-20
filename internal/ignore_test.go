package gotestx

import "testing"

func TestIgnore(t *testing.T) {
	tests := []struct {
		name     string
		pkgs     []string
		patterns []string
		pkg      string
		expect   bool
		expectPk []string
	}{
		{"MatchSingleSegmentRoot", nil, []string{"mock"}, "mock", true, nil},
		{"MatchSingleSegmentNested", nil, []string{"mock"}, "pkg/mock", true, nil},
		{"MatchSingleSegmentDeep", nil, []string{"mock"}, "pkg/a/mock", true, nil},
		{"NoMatchPrefix", nil, []string{"mock"}, "pkg/a/mockery", false, nil},

		{"ExactPathMatch", nil, []string{"outport/testkit"}, "outport/testkit", true, nil},
		{"ExactPathNested", nil, []string{"outport/testkit"}, "pkg/outport/testkit", true, nil},
		{"ExactPathTooDeep", nil, []string{"outport/testkit"}, "outport/x/testkit", false, nil},
		{"ExactPathMismatch", nil, []string{"outport/testkit"}, "pkg/outport/x/testkit", false, nil},

		{"MultiplePatternsFirstMatch", nil, []string{"mock", "testkit"}, "pkg/a/mock", true, nil},
		{"MultiplePatternsSecondMatch", nil, []string{"mock", "testkit"}, "pkg/a/testkit", true, nil},
		{"MultiplePatternsNoMatch", nil, []string{"mock", "testkit"}, "pkg/a/service", false, nil},

		{"LeadingDotSlash", nil, []string{"pkg/service"}, "./pkg/service", true, nil},
		{"TrailingSlash", nil, []string{"pkg/service"}, "pkg/service/", true, nil},
		{"PatternWithSpaces", nil, []string{"  mock  "}, "pkg/mock", true, nil},

		{"EmptyPkg", nil, []string{"mock"}, "", false, nil},
		{"NilPatterns", nil, nil, "pkg/mock", false, nil},
		{"EmptyPatterns", nil, []string{}, "pkg/mock", false, nil},
		{"EmptyPatternEntriesIgnored", nil, []string{"", "   ", "mock"}, "pkg/mock", true, nil},
		{"OnlyEmptyPatterns", nil, []string{"", "   "}, "pkg/mock", false, nil},
		{"BothEmpty", nil, []string{}, "", false, nil},

		{"ExactMatchOnly", nil, []string{"service"}, "pkg/service", true, nil},
		{"NoPartialMatch", nil, []string{"service"}, "pkg/services", false, nil},

		{"DeepMatch", nil, []string{"testkit"}, "a/b/c/testkit", true, nil},
		{"DeepNoMatch", nil, []string{"testkit"}, "a/b/c/testkits", false, nil},

		{"PatternLongerThanPkg", nil, []string{"a/b/c"}, "a/b", false, nil},
		{"NoMatchSlidingWindow", nil, []string{"a/b"}, "x/y/z", false, nil},
		{"MismatchFirstSegment", nil, []string{"a/b"}, "x/b", false, nil},
		{"MismatchSecondSegment", nil, []string{"a/b"}, "a/x", false, nil},

		{"FilterNoPatterns", []string{"pkg/a", "pkg/b"}, nil, "", false, []string{"pkg/a", "pkg/b"}},
		{"FilterEmptyPackages", []string{}, []string{"mock"}, "", false, []string{}},
		{"FilterSingleMatch", []string{"pkg/a", "pkg/mock", "pkg/b"}, []string{"mock"}, "", false, []string{"pkg/a", "pkg/b"}},
		{"FilterMultipleMatches", []string{"pkg/a/mock", "pkg/b/testkit", "pkg/c"}, []string{"mock", "testkit"}, "", false, []string{"pkg/c"}},
		{"FilterExactPath", []string{"pkg/outport/testkit", "pkg/a"}, []string{"outport/testkit"}, "", false, []string{"pkg/a"}},
		{"FilterNoMatch", []string{"pkg/a", "pkg/b"}, []string{"mock"}, "", false, []string{"pkg/a", "pkg/b"}},
		{"FilterAllFiltered", []string{"pkg/mock", "pkg/testkit"}, []string{"mock", "testkit"}, "", false, []string{}},
		{"FilterOrderPreserved", []string{"pkg/a", "pkg/mock", "pkg/b", "pkg/testkit"}, []string{"mock", "testkit"}, "", false, []string{"pkg/a", "pkg/b"}},
		{"FilterPatternWithSpaces", []string{"pkg/mock"}, []string{"  mock  "}, "", false, []string{}},
		{"FilterEmptyPatternEntriesIgnored", []string{"pkg/mock", "pkg/a"}, []string{"", "   ", "mock"}, "", false, []string{"pkg/a"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pkg != "" || tt.patterns != nil {
				got := shouldIgnore(tt.pkg, tt.patterns)
				if got != tt.expect {
					t.Fatalf("shouldIgnore expected %v, got %v (pkg=%q patterns=%v)", tt.expect, got, tt.pkg, tt.patterns)
				}
			}

			if tt.pkgs != nil {
				got := filterPackages(tt.pkgs, tt.patterns)
				if len(got) != len(tt.expectPk) {
					t.Fatalf("filterPackages expected %v, got %v", tt.expectPk, got)
				}
				for i := range got {
					if got[i] != tt.expectPk[i] {
						t.Fatalf("filterPackages expected %v, got %v", tt.expectPk, got)
					}
				}
			}
		})
	}
}
