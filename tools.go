//go:build tools
// +build tools

// This file imports packages that are required when running go generate or used
// as part of the development process, but not depended upon by built code.

package tools

import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
