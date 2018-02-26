package main

//ICommand interface
type ICommand interface {
	run() bool
	getName() string
	build(name string, description string)
	addOption(shortFlag string, longFlag string)
	getBindedInterfaceByFlag(shortFlag string) interface{}
}
