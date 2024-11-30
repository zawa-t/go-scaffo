package project

type Configuration struct {
	Dir   string
	Files map[string]string
}

type Configurations []Configuration

type configuration interface {
	LoadConfigurations(configName string, basePath string, appName string) Configurations
}
