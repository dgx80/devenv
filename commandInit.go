package main

import (
	"github.com/dgx80/goio"
)

type CommandInit struct {
	Force bool
}

func (c *CommandInit, h Help) build() {

}
func (c CommandInit) run() {
	goio.FileExist('')
}