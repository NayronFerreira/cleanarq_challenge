//go:build tools
// +build tools

package tools

import (
	// Import the actual package instead of the program
	_ "github.com/99designs/gqlgen/graphql"
)
