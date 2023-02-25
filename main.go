package main

import (
	"fmt"
	aa "import/a"
	_ "import/b"
)

var name = "main"

func init() {
	fmt.Println(name)
	aa.Test()
}

func main() {
	fmt.Println("exec main")
}
