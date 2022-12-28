//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/playwright-community/playwright-go/cmd/playwright"
)
