package template

import "embed"

type Content struct {
	Dir   string
	Files map[string]string
}

type Config struct {
	EmbededFiles embed.FS
	Contents     []Content
}

type Data struct {
	AppName        string
	BaseImportPath string
}
