# Greetings Module

This is a module for handling greetings in your application.

## Installation

To install the module, run the following command:

```bash
go get github.com/willbipo/greetings
```

## Examples

Here are some examples of how to use the `greetings` module:

### Function ```Hello(name string)```

```go
package main

import (
    "fmt"
    "github.com/willbipo/greetings"
)

func main() {
    name := "John"
    message, err := greetings.Hello(name)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(message)
}
```

Returns:
```
¡Hola, Jose! ¡Bienvenido!
```

### Function ```Hellos(names []string)```

```go
package main

import (
	"fmt"
	"log"
	"github.com/willbipo/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{
		"Jose",
		"Armando",
		"Will",
	}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range message {
		fmt.Println(value)
	}
}
```

Returns:
```
¡Hola, Jose! ¡Bienvenido!
¡Hola, Armando! ¡Bienvenido!
¡Hola, Will! ¡Bienvenido!
```

Powered By Williams Cordova