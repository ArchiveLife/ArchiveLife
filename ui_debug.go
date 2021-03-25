// +build debug

package ui

import (
	"io"
	"io/fs"
)

func GetUIContent() fs.FS {

	if f, err := fs.Sub(content, "build"); err == nil {
		return f
	}
	return nil
}
