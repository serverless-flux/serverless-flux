package main

import (
	"github.com/serverless-operator/serverless-operator/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version will output the build information",
		Long:  "",
		Run: func(_ *cobra.Command, _ []string) {
			config := getConfig()

			logger, err := logging.Configure(config.LoggingConfig)
			if err != nil {
				logrus.Fatalf("failed to configure logging: %s", err.Error())
			}
			logger.Infof("Version: %s", version)
			logger.Infof("Build Date: %s", date)
			logger.Infof("Git Commit: %s", commit)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
