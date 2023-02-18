package main

import (
	"fmt"
	_ "import/a"
	_ "import/b"
)

var name = "main"

func init() {
	fmt.Println(name)
}

func main() {
	fmt.Println("exec main")
}
