package main

import (
	"github.com/serverless-operator/serverless-operator/pkg/logging"
	"github.com/serverless-operator/serverless-operator/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
			logger.Infof("Version: %s", version.Version)
			logger.Infof("Build Date: %s", version.BuildDate)
			logger.Infof("Git Commit: %s", version.GitCommit)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
