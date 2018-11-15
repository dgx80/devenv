package app

import (
	"github.com/dgx80/devenv/dir"
	"github.com/dgx80/devenv/logger"
	"github.com/dgx80/devenv/tree"
)

//DryRun to execute anything on disk
var DryRun bool

// PrepareApp call this function before use any function
func PrepareApp() {
	dir.SetChangePermission(!DryRun)
	if DryRun {
		logger.Verbose = true
	}
}

//Init init folder base
func Init() {
	f := tree.CreateFolder("", false)
	logger.Info("root: " + f.Dev())

	dir.EnsureThatFolderExist(f.Dev())
	dir.EnsureThatFolderExist(f.Scripts())
	dir.EnsureThatFolderExist(f.Projects())
}
