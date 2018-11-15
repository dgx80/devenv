package dir

import (
	"log"
	"os"

	"github.com/dgx80/devenv/logger"
)

var changePermision bool

//SetChangePermission receive permission to change directory
func SetChangePermission(value bool) {
	changePermision = value
}

//EnsureThatFolderExist be sure that folder is created
func EnsureThatFolderExist(fullPath string) {

	logger.Info("EnsureThatFolderExist: " + fullPath)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		logger.Info(fullPath + " not exists")
		if changePermision {
			err = os.Mkdir(fullPath, 0755)
			if err != nil {
				logger.Info("failed")
				log.Fatal(err.Error())
			}
		}
	} else {
		logger.Info(fullPath + " already exists")
	}
}
