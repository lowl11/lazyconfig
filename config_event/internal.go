package config_event

const (
	environmentName    = "env"
	environmentDefault = "test"

	configFolderBase = "config/"
)

type Object[T any] struct {
	Body T
}

func (event *Event[T]) fileName(value string) string {
	return configFolderBase + value + ".json"
}
