package config_map_event

type Event struct {
	object map[string]string
	env    map[string]string
}

func New() *Event {
	return &Event{
		env: make(map[string]string),
	}
}
