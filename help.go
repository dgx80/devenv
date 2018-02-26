package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Help use to send argument
type Help struct {
	ProjectName string
	Force       bool
	Action      string
}

func (h Help) showHelpTitle() {
	fmt.Print(`
USAGE:
	devenv COMMAND [OPTIONS]

		`)
}

func (h Help) showHelp() {
	h.showHelpTitle()
	fmt.Print(`

Commands:
	init		Initialize dev enviroments
	install		Install environments variables .env of project in ~/dev/.env
	uninstall	Uninstall environments variables .env of project in ~/dev/.env
		`)
}

func (h Help) showHelpInstallCommand() {
	h.showHelp()
	fmt.Print(`

Options:
	
	-p, --project-name string		Project name
				`)
}

func (h Help) showHelpInitCommand() {
	h.showHelp()
	fmt.Print(`

Options:
	
	-f, --force		force to create if dev folder exists
				`)
}

func (h *Help) run() {

	if len(os.Args) > 1 {
		h.Action = os.Args[1]
	} else {
		h.Action = ""
	}

	flag.Usage = h.showHelp
	flag.Parse()

	switch h.Action {
	case "init":
		h.handleInitCommand()
	case "install":
		h.handleInstallCommand()
	case "uninstall":
		h.handleUninstallCommand()
	default:
		color.Red("command missing")
		h.showHelp()
		os.Exit(1)
	}
}

func (h Help) handleInitCommand() {
	flag.BoolVar(&h.Force, "f", false, "force")
	flag.BoolVar(&h.Force, "force", false, "force")
}

func (h Help) handleInstallCommand() {
	flag.StringVar(&h.ProjectName, "p", "", "project")
	flag.StringVar(&h.ProjectName, "project-name", "", "project")
	if h.ProjectName == "" {
		h.showHelpInstallCommand()
		os.Exit(1)
	}
}

func (h Help) handleUninstallCommand() {
	//TBD
}
