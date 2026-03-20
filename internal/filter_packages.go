package gotestx

// filterPackages returns packages excluding those matching ignore patterns.
func filterPackages(pkgs []string, patterns []string) []string {
	if len(pkgs) == 0 || len(patterns) == 0 {
		return pkgs
	}

	var filtered []string

	for _, pkg := range pkgs {
		if shouldIgnore(pkg, patterns) {
			continue
		}
		filtered = append(filtered, pkg)
	}

	return filtered
}
