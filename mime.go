package utils

import (
	"mime"
	"path/filepath"
	"strings"
)

var defaultMime = "application/octet-stream"

func ContentType(name string) string {
	ext := filepath.Ext(name)
	if ext == "" {
		return defaultMime
	}

	t := mime.TypeByExtension(ext)
	if t != "" {
		return strings.Split(t, ";")[0]
	}

	return defaultMime
}

func ContentGroup(name string) string {
	t := ContentType(name)
	return strings.Split(t, "/")[0]
}
