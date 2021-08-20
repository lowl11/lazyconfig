# LazyConfig

> Библиотека для генерирования и поддержания конфигурационных файлов для приложений Go

Пример использования
```go
package main

import (
	"github.com/Lowl11/lazyconfig"
	"log"
)

type Configuration struct {
	TestKey string `json:"test_key"`
}

func main() {
	config := &Configuration{}
	debug := true
	if err := lazyconfig.ReadConfig(&config, debug); err != nil {
		log.Fatal(err)
	}	
}
```