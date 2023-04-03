package config_map_event

import "github.com/lowl11/lazyconfig/config_data"

func (event *Event) fileName(value string) string {
	return config_data.ConfigFolder + value + ".yml"
}
