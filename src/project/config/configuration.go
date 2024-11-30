package config

type Structure struct {
	Dir   string
	Files map[string]string
}

type Contents []Structure

type Configuration struct {
	TemplatePath string
	Contents     Contents
}
