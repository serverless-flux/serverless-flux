package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	kubeconfig string
	cfgFile    string
	logLevel   string
	logFile    string

	rootCmd = &cobra.Command{
		Use:   "serverless-operator",
		Short: "Serverless Operator is used to deploy and manage Serverless.com appliations",
		Long: `Serverless Operator is used to deploy and manage Serverless.com appliations.
			With a manifest file you can deploy, update or delete your serverless application.`,
		Run: func(c *cobra.Command, _ []string) {
			_ = c.Help()
		},
	}
)

func main() {
	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config", "f",
		"Config file (default is $HOME/.serverless-operator.yaml)")
	rootCmd.PersistentFlags().StringVarP(&kubeconfig,
		"kubeconfig", "k", "",
		"Absolute path to the kubeconfig file. Only required if out-of-cluster.")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "Info", "Log level for the CLI")
	rootCmd.PersistentFlags().StringVarP(&logFile, "logfile", "", "", "Log level for the CLI")

	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("kubeconfig", rootCmd.PersistentFlags().Lookup("kubeconfig"))
	_ = viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
	_ = viper.BindPFlag("logfile", rootCmd.PersistentFlags().Lookup("logfile"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".serverless-operator")
	}

	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
