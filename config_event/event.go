package config_event

type Event[T any] struct {
	object             Object[T]
	env                map[string]string
	envFileName        string
	environmentName    string
	environmentDefault string
}

func New[T any]() *Event[T] {
	return &Event[T]{
		env:                make(map[string]string),
		envFileName:        envFileNameDefault,
		environmentName:    environmentName,
		environmentDefault: environmentDefault,
	}
}
