package main

import (
	"fmt"
)

var flagvar string

func main() {
	h := Help{}
	h.run()

	fmt.Println(h.Name)
}
