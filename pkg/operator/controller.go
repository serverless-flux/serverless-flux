package operator

import (
	opkit "github.com/rook/operator-kit"
	sls "github.com/serverless-operator/serverless-operator/pkg/apis/serverlessrelease/v1alpha1"

	// nolint:lll
	slsclient "github.com/serverless-operator/serverless-operator/pkg/client/clientset/versioned/typed/serverlessrelease/v1alpha1"
	"github.com/serverless-operator/serverless-operator/pkg/config"
	"github.com/serverless-operator/serverless-operator/pkg/env"
	"k8s.io/client-go/tools/cache"
)

// Controller represents the controller object for the resource
type Controller struct {
	config       *config.Config
	context      *opkit.Context
	slsClientSet slsclient.ReleaseV1alpha1Interface
}

// NewController creates a new Controller
func NewController(config *config.Config,
	context *opkit.Context,
	slsClientSet slsclient.ReleaseV1alpha1Interface) *Controller {
	return &Controller{
		config:       config,
		context:      context,
		slsClientSet: slsClientSet,
	}
}

// StartWatch watches for instances of Object Store custom resources and acts on them
func (c *Controller) StartWatch(namespace string, stopCh <-chan struct{}) error {

	resourceHandlers := cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	}
	restClient := c.slsClientSet.RESTClient()
	watcher := opkit.NewWatcher(Resource, namespace, resourceHandlers, restClient)

	// nolint:errcheck
	go watcher.Watch(&sls.ServerlessRelease{}, stopCh)

	return nil
}

func (c *Controller) onAdd(obj interface{}) {
	slsRelease := obj.(*sls.ServerlessRelease).DeepCopy()
	logger := c.config.Logger

	logger.Infof("Serverless release added %s and package %s\n",
		slsRelease.Spec.ReleaseName,
		slsRelease.Spec.PackagePath)

	ns := slsRelease.Namespace

	for _, envVar := range slsRelease.Spec.Env {
		if envVar.ValueFrom == nil {
			logger.Infof("env variable %s=%s\n", envVar.Name, envVar.Value)
			continue
		}

		value, err := env.GetEnvVarValue(c.config.KubeClientset, ns, envVar.ValueFrom)
		if err != nil {
			logger.Errorf("error getting environment variable: %v", err)
			continue
		}
		logger.Infof("env variable (from ref) %s=%s\n", envVar.Name, value)
	}

}

func (c *Controller) onUpdate(oldObj, newObj interface{}) {
	slsReleaseOld := oldObj.(*sls.ServerlessRelease).DeepCopy()
	slsReleaseNew := newObj.(*sls.ServerlessRelease).DeepCopy()

	c.config.Logger.Infof("Serverless release updated from %s to package %s\n",
		slsReleaseOld.Spec.PackagePath,
		slsReleaseNew.Spec.PackagePath)
}

func (c *Controller) onDelete(obj interface{}) {
	slsRelease := obj.(*sls.ServerlessRelease).DeepCopy()

	c.config.Logger.Infof("Serverless release deleted %s and package %s\n",
		slsRelease.Spec.ReleaseName,
		slsRelease.Spec.PackagePath)
}
