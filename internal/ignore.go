package gotestx

import "strings"

// shouldIgnore determines if a package matches any ignore pattern.
//
// Tree-based matching rules:
//
//   - "mock" matches any segment named "mock"
//   - "outport/testkit" matches exact contiguous path
//   - No globbing, no wildcards, no prefix matching
func shouldIgnore(pkg string, patterns []string) bool {
	if pkg == "" || len(patterns) == 0 {
		return false
	}

	pkgParts := split(normalize(pkg))

	for _, raw := range patterns {
		p := normalize(strings.TrimSpace(raw))
		if p == "" {
			continue
		}

		patternParts := split(p)

		if matchTree(pkgParts, patternParts) {
			return true
		}
	}

	return false
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "./")
	s = strings.Trim(s, "/")
	return s
}

func split(s string) []string {
	return strings.Split(s, "/")
}

// matchTree checks if pattern exists as contiguous sequence in pkg
func matchTree(pkg, pattern []string) bool {
	if len(pattern) == 0 || len(pkg) < len(pattern) {
		return false
	}

	for i := 0; i <= len(pkg)-len(pattern); i++ {
		if matchExact(pkg[i:i+len(pattern)], pattern) {
			return true
		}
	}

	return false
}

func matchExact(a, b []string) bool {
	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
