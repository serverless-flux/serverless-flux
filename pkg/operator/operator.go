package operator

import (
	"fmt"
	"time"

	opkit "github.com/rook/operator-kit"
	slsscheme "github.com/serverless-operator/serverless-operator/pkg/client/clientset/versioned/scheme"
	slsclient "github.com/serverless-operator/serverless-operator/pkg/client/clientset/versioned/typed/serverlessrelease/v1alpha1"
	"github.com/serverless-operator/serverless-operator/pkg/config"
	corev1 "k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
)

const controllerName = "serverless-operator"

// New creates a new server from a config
func New(config *config.Config) *Operator {
	return &Operator{
		Config: config,
	}
}

// Run starts the operator listening for Kuberbetes events
func (o *Operator) Run(stopChan <-chan struct{}) {
	logger := o.Config.Logger
	logger.Info("Getting Kubernetes context")
	context, restConfig, kubeClientset, slsClientSet, err := createContext(o.Config.Kubeconfig)
	if err != nil {
		logger.Fatalf("Failed to create context. %+v\n", err)
	}
	o.Config.Context = context
	o.Config.SlsClientSet = slsClientSet
	o.Config.RESTConfig = restConfig

	logger.Info("Registering resources")
	resources := []opkit.CustomResource{Resource}
	err = opkit.CreateCustomResources(*context, resources)
	if err != nil {
		logger.Fatalf("Failed to create customer resource. %+v\n", err)
	}

	slsscheme.AddToScheme(scheme.Scheme)
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(logger.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerName})
	o.Config.Recorder = recorder

	logger.Info("Watching the resources")
	controller := NewController(o.Config, context, slsClientSet)
	controller.StartWatch(corev1.NamespaceAll, stopChan)
}

func createContext(kubeconfig string) (*opkit.Context, *rest.Config, kubernetes.Interface, slsclient.ReleaseV1alpha1Interface, error) {
	config, err := getClientConfig(kubeconfig)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to get k8s config. %+v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to get k8s client. %+v", err)
	}

	apiExtClientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to create k8s API extension clientset. %+v", err)
	}

	slsclientset, err := slsclient.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to create object store clientset. %+v", err)
	}

	context := &opkit.Context{
		Clientset:             clientset,
		APIExtensionClientset: apiExtClientset,
		Interval:              500 * time.Millisecond,
		Timeout:               60 * time.Second,
	}

	return context, config, clientset, slsclientset, nil

}

func getClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
