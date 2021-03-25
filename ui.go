// +build !debug

package ui

import (
	"embed"
	"io/fs"
)

//go:embed build/*
var content embed.FS

func GetUIContent() fs.FS {
	if f, err := fs.Sub(content, "build"); err == nil {
		return f
	}
	return nil
}
