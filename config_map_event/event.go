package config_map_event

type Event struct {
	object map[string]string

	env                map[string]string
	envFileName        string
	environmentName    string
	environmentDefault string
}

func New() *Event {
	return &Event{
		env:                make(map[string]string),
		envFileName:        envFileNameDefault,
		environmentName:    environmentName,
		environmentDefault: environmentDefault,
	}
}
