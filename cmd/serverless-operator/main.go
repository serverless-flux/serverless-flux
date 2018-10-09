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
	master     string
	cfgFile    string
	logLevel   string
	logFile    string

	rootCmd = &cobra.Command{
		Use:   "serverless-operator",
		Short: "Serverless Operator is used to deploy and manage Serverless.com appliations",
		Long:  "Serverless Operator is used to deploy and manage Serverless.com appliations. With a manifest file you can deploy, update or delete your serverless application.",
		Run: func(c *cobra.Command, _ []string) {
			c.Help()
		},
	}
)

func main() {
	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
		flag.CommandLine.Parse([]string{"-logtostderr"})
		fs.Parse(os.Args)

		//TODO: output version

		{
			logger = log.NewLogfmtLogger(os.Stderr)
			logger = log.With(logger, "ts", log.DefaultTimestampUTC)
			logger = log.With(logger, "caller", log.DefaultCaller)
		}

		errc := make(chan error)

		// Shutdown trigger for goroutines
		shutdown := make(chan struct{})
		shutdownWg := &sync.WaitGroup{}

		go func() {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			errc <- fmt.Errorf("%s", <-c)
		}()

		defer func() {
			logger.Log("exiting...", <-errc)
			close(shutdown)
			shutdownWg.Wait()
		}()

		clusterConfig, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)
		if err != nil {
			//TODO: Log error
		}

		kubeClient, err := kubernetes.NewForConfig(clusterConfig)
		if err != nil {
			//TODO: log error
		}
	*/

}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "f", "Config file (default is $HOME/.serverless-operator.yaml)")
	rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "Absolute path to the kubeconfig file. Only required if out-of-cluster.")
	rootCmd.PersistentFlags().StringVarP(&master, "master", "m", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "Info", "Log level for the CLI")
	rootCmd.PersistentFlags().StringVarP(&logFile, "logfile", "", "", "Log level for the CLI")

	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("kubeconfig", rootCmd.PersistentFlags().Lookup("kubeconfig"))
	viper.BindPFlag("master", rootCmd.PersistentFlags().Lookup("master"))
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
	viper.BindPFlag("logfile", rootCmd.PersistentFlags().Lookup("logfile"))
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
