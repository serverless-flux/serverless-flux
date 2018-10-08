package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/spf13/pflag"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	logger log.Logger
	fs     *pflag.FlagSet

	kubeconfig *string
	master     *string
)

func main() {
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

}

func init() {
	fs = pflag.NewFlagSet("default", pflag.ExitOnError)

	kubeconfig = fs.String("kubeconfig", "", "Absolute path to the kubeconfig file. Only required if out-of-cluster.")
	master = fs.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")

}
