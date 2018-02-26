package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
)

//CommandManager contain all command in project
type CommandManager struct {
	commands map[string]ICommand
	help     Help
}

func (cm *CommandManager) init() {
	cm.help = Help{}
	cm.commands = make(map[string]ICommand)
}

func (cm *CommandManager) addCommand(c ICommand) {
	s := c.getName()
	cm.commands[s] = c
}

func (cm CommandManager) findCommand() ICommand {
	commandName := ""

	if len(os.Args) > 1 {
		commandName = os.Args[1]
	} else {
		color.Red("command missing")
		cm.help.showHelp()
		os.Exit(1)
	}

	flag.Parse()

	command, exists := cm.commands[commandName]

	if exists {
		color.Green(commandName)
	} else {
		color.Red("command is wrong")
		cm.help.showHelp()
		os.Exit(1)
	}
	return command
}

func (cm CommandManager) run() bool {
	command := cm.findCommand()

	command.run()
	return true
}
