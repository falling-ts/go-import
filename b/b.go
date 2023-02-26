package b

import (
	"fmt"
	_ "github.com/falling-ts/import/c"
	_ "github.com/falling-ts/import/h"
)

var name = "b"

func init() {
	fmt.Println(name)
}
