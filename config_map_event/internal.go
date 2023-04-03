package config_map_event

const (
	environmentName    = "env"
	environmentDefault = "test"

	configFolderBase   = "profiles/"
	envFileNameDefault = ".env"
)

func (event *Event) fileName(value string) string {
	return configFolderBase + value + ".yml"
}
