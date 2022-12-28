//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/playwright-community/playwright-go/cmd/playwright"
)
