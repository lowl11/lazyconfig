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

#### Example with all methods
```go
config, err := confapi.New[Configuration]().
		EnvironmentDefault("local"). // if there is no any value in "env" variable, you can set default
		EnvironmentName("").         // usually lib looking for "env" variable, but you can change it
		EnvFileName("").             // for replace variables usually uses ".env" file, but you can change it
		Read()
if err != nil {
    log.Fatal("Read config err:", err)
}
```