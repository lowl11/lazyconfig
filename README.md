# lazyconfig

> Simple config library with big customization

### Get started
Create directory <b>/profiles</b> in the root of your project <br>
Configuration file name depends on your "env" variable <br>
If "env" variables equals to "production", configuration file name should be "production.yml" <br>
For not duplicating variables in configuration files, you can create base <b>config.yml</b> file 

First, need to initialize config
```go
import (
    "github.com/lowl11/lazyconfig/config/config_internal"
)

func main() {
    config_internal.Init()
}
```

After that, you have access to config variables from anywhere
```go
import (
    "fmt"
    "github.com/lowl11/lazyconfig/config"
    "github.com/lowl11/lazyconfig/config/config_internal"
)

func main() {
    // initialize config variables
    config_internal.Init()
	
    // get variables 
    fmt.Println("database connection:", config.Get("database_connection"))
    fmt.Println("max connections:", config.Get("max_connections"))
    fmt.Println("some basic value:", config.Get("some_basic_thing"))
    fmt.Println("colvir username:", config.Get("colvir_username"))
    fmt.Println("colvir password:", config.Get("colvir_password"))	
}
```

### Customization
