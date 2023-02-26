package aa

import (
	"fmt"
	jj "github.com/falling-ts/import/a/j"
	_ "github.com/falling-ts/import/c"
	_ "github.com/falling-ts/import/d"
)

var name = "a"

func init() {
	fmt.Println(name)
	jj.Test()
}

func Test() {
	fmt.Println("aa")
}
