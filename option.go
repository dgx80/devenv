package main

//Option to add to command
type Option struct {
	v         interface{}
	shortFlag string
	longFlag  string
}

func (o *Option) create(shortFlag string, longFlag string) {
	o.shortFlag = shortFlag
	o.longFlag = longFlag
}
