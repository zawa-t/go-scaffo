package config

import "embed"

type Content struct {
	Dir   string
	Files map[string]string
}

type Template struct {
	EmbededFiles embed.FS
	Path         string
}

type Configuration struct {
	Template Template
	Contents []Content
}
