package main

//CommandInit to intitialize dev folder
type CommandInit struct {
	Force bool
	base  CommandBase
}

func (c *CommandInit) build(name string, description string) {
	c.base = CommandBase{}
	c.base.Name = name
	c.base.Description = description
}

func (c CommandInit) run() bool {
	return false
}

func (c *CommandInit) addOption(shortFlag string, longFlag string) {
	c.base.addOption(shortFlag, longFlag)
}

func (c CommandInit) getName() string {
	return c.base.Name
}

func (c *CommandInit) getBindedInterfaceByFlag(shortFlag string) interface{} {
	switch shortFlag {
	case "f":
		return c.Force
	}
	return nil
}
