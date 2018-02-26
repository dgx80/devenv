package main

//CommandBase is the base composition of Command
type CommandBase struct {
	Name        string
	Description string
	options     []Option
}

func (c *CommandBase) addOption(shortFlag string, longFlag string) {
	o := Option{}
	o.create(shortFlag, longFlag)
	c.options = append(c.options, o)
}
