package cmd

import (
	"fmt"

	"os"

	"github.com/a-berahman/gitpipe/utility/logger"
	"github.com/spf13/cobra"
)

var configPath string

var rootCMD = &cobra.Command{
	Use:   "gitpipe",
	Short: "gitpipe handler of GitHub API and Pipedrive API",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCMD.PersistentFlags().StringVarP(&configPath, "config-path", "c", "env.yaml", "path to config directory")

}

func initConfig() {
	logger.Initialize()
}

//Execute runs through the command tree finding appropriate matches for commands and then corresponding flags
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
