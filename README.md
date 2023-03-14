# lazyconfig

> Simple config library with big customization

#### Simple example:
```go
package main

import (
	"github.com/lowl11/lazyconfig"
	"log"
)

type Configuration struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}

func main() {
	config, err := confapi.New[Configuration]().
		EnvironmentDefault("local").
		Read()
	if err != nil {
		log.Fatal("Read config err:", err)
	}

	fmt.Printf("Config: %v\n", config)
}
```