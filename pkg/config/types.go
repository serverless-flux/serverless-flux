package config

import (
	opkit "github.com/rook/operator-kit"
	// nolint:lll
	slsclient "github.com/serverless-operator/serverless-operator/pkg/client/clientset/versioned/typed/serverlessrelease/v1alpha1"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
)

// Config represent the configuration of the operator
type Config struct {
	Kubeconfig     string
	KubeClientset  kubernetes.Interface
	SlsClientSet   slsclient.ReleaseV1alpha1Interface
	RESTConfig     *rest.Config
	Context        *opkit.Context
	LoggingConfig  *LoggingConfig
	Logger         *logrus.Entry
	Depenendencies *DependenciesConfig
	//TODO: add the rest
	Recorder record.EventRecorder
}

// LoggingConfig defines the attributes for the logger
type LoggingConfig struct {
	File              string
	Level             string
	DisableTimestamps bool
	FullTimestamps    bool
}

// DependenciesConfig defines the config relevenat to dependencies
type DependenciesConfig struct {
	ServerlessPath string
	NodePath       string
}
