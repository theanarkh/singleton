# singleton
this is a singleton pattern implementation in golang.
# Usage
```go
package main
import (
	"fmt"
	"github.com/theanarkh/singleton"
)
type Dummy struct {}

func factory() (*Dummy, error) {
	return &Dummy{}, nil
}

func main() {
	singleton := singleton.New(factory)
	instance, err := singleton.Get()
	fmt.Println(instance, err)
}
```