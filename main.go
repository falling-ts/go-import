package main

import (
	"fmt"
	aa "github.com/falling-ts/import/a"
	_ "github.com/falling-ts/import/b"
	_ "github.com/falling-ts/import/c"
	_ "github.com/falling-ts/import/h"
)

var name = "main"

func init() {
	fmt.Println(name)
	aa.Test()
}

func main() {
	fmt.Println("exec main")
}
