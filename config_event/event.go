package config_event

type Event[T any] struct {
	object             Object[T]
	configurations     map[string]string
	environmentName    string
	environmentDefault string
}

func New[T any]() *Event[T] {
	return &Event[T]{
		configurations: map[string]string{
			"local":      "local",
			"test":       "test",
			"production": "production",
		},
		environmentName:    environmentName,
		environmentDefault: environmentDefault,
	}
}
