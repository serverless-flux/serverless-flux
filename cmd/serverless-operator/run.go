package main

import (
	"github.com/serverless-operator/serverless-operator/pkg/config"
	"github.com/serverless-operator/serverless-operator/pkg/dependencies"
	"github.com/serverless-operator/serverless-operator/pkg/logging"
	"github.com/serverless-operator/serverless-operator/pkg/operator"
	"github.com/serverless-operator/serverless-operator/pkg/signal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the operator",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config := getConfig()

		logger, err := logging.Configure(config.LoggingConfig)
		if err != nil {
			logrus.Fatalf("failed to configure logging: %s", err.Error())
		}
		config.Logger = logger

		err = dependencies.Check(config)
		if err != nil {
			logger.Fatalf("dependency check failed: %s", err.Error())
		}

		stopChan := signal.SetupSignalHandler()

		operator.New(config).Run(stopChan)

		logger.Info("started operator")
		<-stopChan
		logger.Info("shutdown signal received, exiting...")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func getConfig() *config.Config {
	config := &config.Config{
		Kubeconfig: kubeconfig,
		LoggingConfig: &config.LoggingConfig{
			File:              logFile,
			Level:             logLevel,
			FullTimestamps:    true,
			DisableTimestamps: false,
		},
		Depenendencies: &config.DependenciesConfig{
			ServerlessPath: serverlessPath,
			NodePath:       nodePath,
		},
	}

	return config
}
