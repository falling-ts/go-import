package e

import (
	"fmt"
	_ "github.com/falling-ts/import/i"
)

var name = "e/1"

var Test = test

func init() {
	fmt.Println(name)
	fmt.Println(Test)
}
