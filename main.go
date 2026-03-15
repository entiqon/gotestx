/*
Copyright © 2025-2026 Entiqon

Author: Isidro A. López G.

This file is part of the GoTestX project.

Repository:
https://github.com/entiqon/gotestx

Licensed under the MIT License.
See the LICENSE file in the project root for license information.
*/
package main

import (
	"os"

	"github.com/entiqon/gotestx/internal"
)

func main() {
	code := gotestx.Run(os.Args[1:], os.Stdout, os.Stderr)
	os.Exit(code)
}
