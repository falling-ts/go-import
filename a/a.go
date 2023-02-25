package aa

import (
	"fmt"
	jj "import/a/j"
	_ "import/c"
	_ "import/d"
)

var name = "a"

func init() {
	fmt.Println(name)
	jj.Test()
}

func Test() {
	fmt.Println("aa")
}
