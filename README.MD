## modulized
```bash
go help mod 
go help mod init
```
```go
package main

import (
	"fmt"

	"github.com/leson/go_poc/greetings"
)

func main() {
	fmt.Println("invoke self-define module of go!")
	fmt.Println(greetings.Hello("Leson"))
}
```

```bash
go mod init go-mod-test
go mod tidy
go run main.go
```

> Go code is grouped into packages, and packages are grouped into modules


## Reference
- [Official Web Site](https://go.dev/)
- [doc of quickstart](https://go.dev/doc/tutorial/getting-started)
- [doc of create-module](https://go.dev/doc/tutorial/create-module)