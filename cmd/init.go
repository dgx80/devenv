// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/dgx80/devenv/app"
	"github.com/dgx80/devenv/logger"
	"github.com/spf13/cobra"
)

var Force bool
var Root string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "To init dev environment",
	Long: `To init dev environment:

It will create dev folder at home.`,
	Run: func(cmd *cobra.Command, args []string) {
		initCmdRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().BoolVarP(&Force, "force", "f", false, "To force initialisation when folder is already created")
}

func initCmdRun(cmd *cobra.Command, args []string) {
	logger.Info("init called")
	app.PrepareApp()
	app.Init()
}
