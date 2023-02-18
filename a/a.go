package a

import (
	"fmt"
	_ "import/c"
	_ "import/d"
)

var name = "a"

func init() {
	fmt.Println(name)
}
